package kafka

import (
	"git.tenvine.cn/backend/gore/log"
	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
)

var (
	ErrInvalidAssignor = errors.New("invalid assignor")
	ErrDisable         = errors.New("kafka is disable")
)

type Config struct {
	Enable bool

	// x.x.x
	Version  string
	Assignor string
	Oldest   bool

	Consumer  ConsumerConfig
	Consumers map[string]ConsumerConfig
}

func NewKafkaConfig(cfg Config) (*sarama.Config, error) {

	if !cfg.Enable {
		return nil, ErrDisable
	}

	sarama.Logger = log.StandardLogger()

	version, err := sarama.ParseKafkaVersion(cfg.Version)
	if err != nil {
		return nil, err
	}

	config := sarama.NewConfig()
	config.Version = version

	if cfg.Assignor != "" {
		switch cfg.Assignor {
		case "sticky":
			config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
		case "roundrobin":
			config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
		case "range":
			config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
		default:
			return nil, errors.Wrap(ErrInvalidAssignor, cfg.Assignor)
		}
	}

	if cfg.Oldest {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	return config, nil
}
