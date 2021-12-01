package kafka

import (
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

	Consumer  ConsumerConfig
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
	//viper := gonfig.Instance()
	//cfg := &Config{
	//	Enable:   viper.GetBool("gore.kafka.enable"),
	//	Version:  viper.GetString("gore.kafka.version"),
	//	Assignor: viper.GetString("gore.kafka.assignor"),
	//	Oldest:   viper.GetBool("gore.kafka.oldest"),
	//	Consumer: ConsumerConfig{
	//		Brokers: viper.GetStringSlice("gore.kafka.consumer.brokers"),
	//		Topics:  viper.GetStringSlice("gore.kafka.consumer.topics"),
	//		Group:   viper.GetString("gore.kafka.consumer.group"),
	//	},
	//}
	//consumers := viper.GetStringMap("gore.kafka.consumers")
	//if len(consumers) == 0 {
	//	return cfg
	//}
	//cfg.Consumers = make(map[string]ConsumerConfig)
	//for key, consumer := range consumers {
	//	b, err := json.Marshal(consumer)
	//	if err != nil {
	//		continue
	//	}
	//	c := ConsumerConfig{}
	//	if err := json.Unmarshal(b, &c); err != nil {
	//		continue
	//	}
	//	cfg.Consumers[key] = c
	//}
	//return cfg
}

func NewKafkaConfig(cfg *Config) (*sarama.Config, error) {

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
	config.ClientID = util.GetLocalhost()

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
