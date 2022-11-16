package kafka

import (
	"fmt"
	"git.tenvine.cn/backend/gore/gonfig"
	"git.tenvine.cn/backend/gore/log"
	"git.tenvine.cn/backend/gore/util"
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

	Consumers map[string]ConsumerConfig
}

type ConsumerConfig struct {
	Brokers []string
	Topics  []string
	Group   string
}

func NewConfig() *Config {
	cfg := Config{}
	if err := gonfig.Instance().UnmarshalKey("gore.kafka", &cfg); err != nil {
		return nil
	}
	return &cfg
}

func NewKafkaConfig(cfg *Config) (*sarama.Config, error) {

	if !cfg.Enable {
		return nil, ErrDisable
	}

	sarama.Logger = log.Default()

	version, err := sarama.ParseKafkaVersion(cfg.Version)
	if err != nil {
		return nil, err
	}

	config := sarama.NewConfig()
	config.Version = version
	config.ClientID = fmt.Sprintf("sarama-%s-%s-%s", gonfig.Instance().GetString("env"), gonfig.Instance().GetString("name"), util.GetLocalhost())

	if len(cfg.Assignor) > 0 {
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
