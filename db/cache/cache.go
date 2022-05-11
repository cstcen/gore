package cache

import (
	"git.tenvine.cn/backend/gore/gonfig"
	"git.tenvine.cn/backend/gore/log"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

var (
	cc *cache.Cache
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

func Setup() error {

	cfg := NewConfig()

	if !cfg.Enable {
		return nil
	}

	options := newOptions(cfg)
	options.LocalCache = cache.NewTinyLFU(1000, time.Minute)
	options.StatsEnabled = !cfg.DisableStats
	cc = cache.New(options)

	return nil
}

func NewConfig() *Config {
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
		log.StandardLogger().Infof("Current redis ring: %+v", addrs)

		options.Redis = redis.NewRing(&redis.RingOptions{
			Addrs:    addrs,
			Username: cfg.Username,
			Password: cfg.Password,
		})
	} else {
		log.StandardLogger().Infof("Current redis cluster: %+v", cfg.Hosts)
		options.Redis = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    cfg.Hosts,
			Username: cfg.Username,
			Password: cfg.Password,
		})
	}
	return options
}
