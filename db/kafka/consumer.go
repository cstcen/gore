package kafka

import (
	"context"
	"git.tenvine.cn/backend/gore/log"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
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
		log.WithFields(logrus.Fields{
			"topic":        message.Topic,
			"timestamp":    message.Timestamp,
			"consumerName": c.name,
		}).Infof(
			"message claimed: %s",
			string(message.Value),
		)
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
	fieldLog := log.WithField("consumerName", c.name)

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
				fieldLog.Warnf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}

			c.ready = make(chan bool)
		}
	}()

	go func() {
		<-c.ready
		fieldLog.Println(c.name + " consumer up and running!...")
		select {
		case <-ctx.Done():
			fieldLog.Infof(c.name + ": context cancelled")
		}

		cancel()
		wg.Wait()

		if err := client.Close(); err != nil {
			fieldLog.Panicf("Error closing client: %v", err)
		}
	}()

	return nil
}

func (c *Consumer) Cancel() {
	c.cancel()
}
