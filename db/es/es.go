package es

import (
	"git.tenvine.cn/backend/gore/gonfig"
	"git.tenvine.cn/backend/gore/log"
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

func Setup() error {
	cfg := NewConfig()

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
	elastic.SetInfoLog(log.StandardLogger())
	elastic.SetTraceLog(log.StandardLogger())
	elastic.SetErrorLog(log.StandardLogger())
	es, err = elastic.NewClientFromConfig(c)
	if err != nil {
		return err
	}

	return nil
}

func NewConfig() *Config {
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
