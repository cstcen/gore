package redis

import (
	"git.tenvine.cn/backend/gore/gonfig"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

var (
	ErrEmptyHosts = errors.New("hosts is empty")

	cli redis.UniversalClient
)

type Config struct {
	Enable         bool
	DisableCluster bool
	Hosts          []string
	Username       string
	Password       string
}

func Instance() redis.UniversalClient {
	return cli
}

func Setup() error {
	cfg := NewConfig()

	if !cfg.Enable {
		return nil
	}

	if len(cfg.Hosts) == 0 {
		return ErrEmptyHosts
	}

	if !cfg.DisableCluster {
		cli = redis.NewClient(&redis.Options{
			Addr:     cfg.Hosts[0],
			Username: cfg.Username,
			Password: cfg.Password,
		})
		return nil
	}

	cli = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    cfg.Hosts,
		Username: cfg.Username,
		Password: cfg.Password,
	})

	return nil
}

func NewConfig() *Config {
	viper := gonfig.Instance()
	cfg := &Config{
		Enable:         viper.GetBool("gore.redis.enable"),
		DisableCluster: viper.GetBool("gore.redis.disableCluster"),
		Hosts:          viper.GetStringSlice("gore.redis.hosts"),
		Username:       viper.GetString("gore.redis.username"),
		Password:       viper.GetString("gore.redis.password"),
	}
	return cfg
}
