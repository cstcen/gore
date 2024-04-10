package kafka

import (
	"github.com/pkg/errors"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

var (
	ErrConsumerConfigNotFound = errors.New("consumer config not found")
	cancellers                = make([]Canceller, 0)
)

func StartupConsumers(handlers map[string]ConsumerMessageHandler) error {
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
		csm := &Consumer{
			ready:         make(chan bool),
			name:          key,
			HandleMessage: handler,
		}
		cancellers = append(cancellers, csm)
		if err := csm.Consume(consumerConfig, config); err != nil {
			return errors.Wrap(err, key)
		}
	}

	return nil
}

func ListeningSigterm() error {
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigterm:
		slog.Info("terminating: via signal")
	}

	if err := ShutdownConsumers(); err != nil {
		return err
	}
	return nil
}

func ShutdownConsumers() error {
	for _, canceller := range cancellers {
		canceller.Cancel()
	}
	return nil
}
