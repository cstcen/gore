package es

import (
	"github.com/cstcen/gore/gonfig"
	"github.com/cstcen/gore/log"
	"github.com/olivere/elastic"
	"github.com/olivere/elastic/config"
)

var (
	es *elastic.Client
)

type Config struct {
	Enable bool
	config.Config
}

func Instance() *elastic.Client {
	return es
}

func SetupDefault() error {
	cfg := DefaultConfig()
	return Setup(cfg)
}

func Setup(cfg *Config) error {

	if !cfg.Enable {
		return nil
	}

	var err error
	c := &config.Config{
		URL:         cfg.URL,
		Index:       cfg.Index,
		Username:    cfg.Username,
		Password:    cfg.Password,
		Shards:      cfg.Shards,
		Replicas:    cfg.Replicas,
		Sniff:       cfg.Sniff,
		Healthcheck: cfg.Healthcheck,
	}
	elastic.SetInfoLog(log.Default())
	elastic.SetTraceLog(log.Default())
	elastic.SetErrorLog(log.Default())
	es, err = elastic.NewClientFromConfig(c)
	if err != nil {
		return err
	}

	return nil
}

func DefaultConfig() *Config {
	viper := gonfig.Instance()
	cfg := &Config{
		Enable: viper.GetBool("gore.elasticsearch.enable"),
		Config: config.Config{
			URL:      viper.GetString("gore.elasticsearch.url"),
			Username: viper.GetString("gore.elasticsearch.username"),
			Password: viper.GetString("gore.elasticsearch.password"),
		},
	}
	return cfg
}
