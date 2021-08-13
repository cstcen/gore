package kafka

import (
	"github.com/pkg/errors"
)

var (
	ErrConsumerConfigNotFound = errors.New("consumer config not found")
)

func SetupConsumers(handlers map[string]ConsumerMessageHandler) error {
	cfg := NewConfig()

	config, err := NewKafkaConfig(cfg)
	if err != nil {
		return err
	}

	for key, handler := range handlers {
		consumerConfig, ok := cfg.Consumers[key]
		if !ok {
			return errors.Wrap(ErrConsumerConfigNotFound, key)
		}
		err := SetupConsumer(key, consumerConfig, config, handler)
		if err != nil {
			return errors.Wrap(err, key)
		}
	}

	return nil
}
