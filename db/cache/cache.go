package cache

import (
	"github.com/cstcen/gore/gonfig"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"log/slog"
	"strconv"
	"time"
)

var (
	cc  *cache.Cache
	cli redis.UniversalClient
)

type Config struct {
	Enable       bool
	EnableRing   bool
	DisableStats bool
	AppName      string
	Hosts        []string
	Username     string
	Password     string
}

func Instance() *cache.Cache {
	return cc
}

func Redis() redis.UniversalClient {
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

	options := newOptions(cfg)
	options.LocalCache = cache.NewTinyLFU(1000, time.Minute)
	options.StatsEnabled = !cfg.DisableStats
	cc = cache.New(options)

	return nil
}

func DefaultConfig() *Config {
	viper := gonfig.Instance()
	cfg := &Config{
		Enable:       viper.GetBool("gore.cache.enable"),
		EnableRing:   viper.GetBool("gore.cache.enableRing"),
		DisableStats: viper.GetBool("gore.cache.disableStats"),
		AppName:      viper.GetString("gore.cache.appName"),
		Hosts:        viper.GetStringSlice("gore.cache.hosts"),
		Username:     viper.GetString("gore.cache.username"),
		Password:     viper.GetString("gore.cache.password"),
	}
	return cfg
}

func newOptions(cfg *Config) *cache.Options {
	options := new(cache.Options)
	if cfg.EnableRing {
		addrs := make(map[string]string)
		for i, host := range cfg.Hosts {
			addrs[cfg.AppName+strconv.Itoa(i)] = host
		}
		slog.Info("Current redis", "ring", addrs)

		cli = redis.NewRing(&redis.RingOptions{
			Addrs:    addrs,
			Username: cfg.Username,
			Password: cfg.Password,
		})
	} else {
		slog.Info("Current redis", "cluster", cfg.Hosts)
		cli = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    cfg.Hosts,
			Username: cfg.Username,
			Password: cfg.Password,
		})
	}
	options.Redis = cli
	return options
}
