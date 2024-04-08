package redis

import (
	"errors"
	"github.com/cstcen/gore/gonfig"
	"github.com/go-redis/redis/v8"
)

const (
	ErrNil = redis.Nil
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

func SetupDefault() error {
	cfg := DefaultConfig()
	return Setup(cfg)
}

func Setup(cfg *Config) error {

	if !cfg.Enable {
		return nil
	}

	if len(cfg.Hosts) == 0 {
		return ErrEmptyHosts
	}

	if !cfg.DisableCluster {
		cli = NewClient(cfg)
		return nil
	}

	cli = NewClusterClient(cfg)

	return nil
}

func NewClusterClient(cfg *Config) *redis.ClusterClient {
	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    cfg.Hosts,
		Username: cfg.Username,
		Password: cfg.Password,
	})
}

func NewClient(cfg *Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Hosts[0],
		Username: cfg.Username,
		Password: cfg.Password,
	})
}

func DefaultConfig() *Config {
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
