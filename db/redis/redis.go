package redis

import (
	"github.com/go-redis/redis/v8"
)

var (
	cli        *redis.Client
	clusterCli *redis.ClusterClient
)

type Config struct {
	Enable bool

	EnableCluster bool
	Hosts         []string
}

func Setup() {
	redis.NewClient(&redis.Options{
		Addr: "",
	})
}
