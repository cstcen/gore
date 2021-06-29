package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

var (
	ErrEmptyHosts = errors.New("hosts is empty")

	cli redis.UniversalClient
)

type Config struct {
	Enable bool

	DisableCluster bool `yaml:"disable-cluster"`
	Hosts          []string
	Username       string
	Password       string
}

func GetInstance() redis.UniversalClient {
	return cli
}

func Setup(cfg Config) error {
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
