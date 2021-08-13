package kafka

import (
	"context"
	"encoding/json"
	"git.tenvine.cn/backend/gore/constant"
	"git.tenvine.cn/backend/gore/gonfig"
	"git.tenvine.cn/backend/gore/log"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type ConsumerMessageHandler interface {
	// HandleConsumerMessage 在ConsumeClaim中调用，用作处理接收到的消息
	HandleConsumerMessage(*sarama.ConsumerMessage)
}

type ConsumerConfig struct {
	Brokers []string
	Topics  []string
	Group   string
}

func StartConsumer(handler ConsumerMessageHandler) error {
	cfg := NewConfig()

	config, err := NewKafkaConfig(cfg)
	if err != nil {
		return err
	}
	return SetupConsumer("gore", cfg.Consumer, config, handler)
}

func NewConfig() *Config {
	viper := gonfig.Instance()
	cfg := &Config{
		Enable:   viper.GetBool("gore.kafka.enable"),
		Version:  viper.GetString("gore.kafka.version"),
		Assignor: viper.GetString("gore.kafka.assignor"),
		Oldest:   viper.GetBool("gore.kafka.oldest"),
		Consumer: ConsumerConfig{
			Brokers: viper.GetStringSlice("gore.kafka.consumer.brokers"),
			Topics:  viper.GetStringSlice("gore.kafka.consumer.topics"),
			Group:   viper.GetString("gore.kafka.consumer.group"),
		},
	}
	consumers := viper.GetStringMap("gore.kafka.consumers")
	if len(consumers) == 0 {
		return cfg
	}
	cfg.Consumers = make(map[string]ConsumerConfig)
	for key, consumer := range consumers {
		b, err := json.Marshal(consumer)
		if err != nil {
			continue
		}
		c := ConsumerConfig{}
		if err := json.Unmarshal(b, &c); err != nil {
			continue
		}
		cfg.Consumers[key] = c
	}
	return cfg
}

func SetupConsumer(name string, cfg ConsumerConfig, config *sarama.Config, handler ConsumerMessageHandler) error {
	fieldLog := log.WithField("consumerName", name)

	consumer := Consumer{name: name, ready: make(chan bool), HandleMessage: handler}

	ctx, cancelFunc := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(cfg.Brokers, cfg.Group, config)
	if err != nil {
		return err
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := client.Consume(ctx, cfg.Topics, &consumer); err != nil {
				fieldLog.Warnf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}

			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready
	fieldLog.Println("Sarama consumer up and running!...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		fieldLog.Println("terminating: context cancelled")
	case <-sigterm:
		fieldLog.Println("terminating: via signal")
	}

	cancelFunc()
	wg.Wait()

	if err = client.Close(); err != nil {
		fieldLog.Warnf("Error closing client: %v", err)
	}

	return err
}

type Consumer struct {
	ready         chan bool
	name          string
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
			"timestamp":    message.Timestamp.Format(constant.FormatTimestamp),
			"consumerName": c.name,
		}).Printf(
			"message claimed: %s",
			string(message.Value),
		)
		session.MarkMessage(message, "")

		c.HandleMessage.HandleConsumerMessage(message)
	}

	return nil
}
