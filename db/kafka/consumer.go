package kafka

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/cstcen/gore/log"
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
		logPrefix := fmt.Sprintf("topic: %s, timestamp: %v, consumer: %s", message.Topic, message.Timestamp, c.name)
		log.Infof("[%s] claimed: %s", logPrefix, string(message.Value))
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
	logPrefix := fmt.Sprintf("consumer: %s", c.name)

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
				log.Warningf("[%s] Error from consumer: %v", logPrefix, err)
			}
			if ctx.Err() != nil {
				return
			}

			c.ready = make(chan bool)
		}
	}()

	go func() {
		<-c.ready
		log.Infof("[%s] consumer up and running!...", logPrefix)
		select {
		case <-ctx.Done():
			log.Infof("[%s] context cancelled", logPrefix)
		}

		cancel()
		wg.Wait()

		if err := client.Close(); err != nil {
			log.Panicf("Error closing client: %v", err)
		}
	}()

	return nil
}

func (c *Consumer) Cancel() {
	c.cancel()
}
