package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"log/slog"
	"sync"
)

type ConsumerMessageHandler interface {
	// HandleConsumerMessage 在ConsumeClaim中调用，用作处理接收到的消息
	HandleConsumerMessage(*sarama.ConsumerMessage) error
}

type Canceller interface {
	Cancel()
}

type Consumer struct {
	ready         chan bool
	name          string
	cancel        context.CancelFunc
	HandleMessage ConsumerMessageHandler
}

func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(c.ready)
	return nil
}

func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		slog.Info("claimed", "topic", message.Topic, "timestamp", message.Timestamp, "consumer", c.name, "value", string(message.Value))
		if c.HandleMessage != nil {
			if err := c.HandleMessage.HandleConsumerMessage(message); err != nil {
				return err
			}
		}
		session.MarkMessage(message, "")
	}

	return nil
}

func (c *Consumer) Consume(cfg ConsumerConfig, config *sarama.Config) error {
	ctx, cancel := context.WithCancel(context.Background())
	c.cancel = cancel
	client, err := sarama.NewConsumerGroup(cfg.Brokers, cfg.Group, config)
	if err != nil {
		cancel()
		return err
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := client.Consume(ctx, cfg.Topics, c); err != nil {
				slog.Warn("consume failed", "consumer", c.name, "err", err)
			}
			if ctx.Err() != nil {
				return
			}

			c.ready = make(chan bool)
		}
	}()

	go func() {
		<-c.ready
		slog.Info("consumer up and running!...", "consumer", c.name)
		select {
		case <-ctx.Done():
			slog.Info("context cancelled", "consumer", c.name)
		}

		cancel()
		wg.Wait()

		if err := client.Close(); err != nil {
			slog.Warn("closing client failed", "err", err)
		}
	}()

	return nil
}

func (c *Consumer) Cancel() {
	c.cancel()
}
