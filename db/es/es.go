package es

import (
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

func GetInstance() *elastic.Client {
	return es
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
	elastic.SetInfoLog(log.StandardLogger())
	elastic.SetTraceLog(log.StandardLogger())
	elastic.SetErrorLog(log.StandardLogger())
	es, err = elastic.NewClientFromConfig(c)
	if err != nil {
		return err
	}

	return nil
}
