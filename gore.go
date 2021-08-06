package gore

import (
	"context"
	"database/sql"
	"git.tenvine.cn/backend/gore/cmd"
	"git.tenvine.cn/backend/gore/consul"
	goreCache "git.tenvine.cn/backend/gore/db/cache"
	goreEs "git.tenvine.cn/backend/gore/db/es"
	"git.tenvine.cn/backend/gore/db/kafka"
	goreKafka "git.tenvine.cn/backend/gore/db/kafka"
	goreMongo "git.tenvine.cn/backend/gore/db/mongo"
	goreMysql "git.tenvine.cn/backend/gore/db/mysql"
	goreRedis "git.tenvine.cn/backend/gore/db/redis"
	goreGin "git.tenvine.cn/backend/gore/gin"
	"git.tenvine.cn/backend/gore/gonfig"
	goreHttp "git.tenvine.cn/backend/gore/http"
	"git.tenvine.cn/backend/gore/infratoken"
	"git.tenvine.cn/backend/gore/log"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func Setup() error {

	if err := gonfig.Setup(); err != nil {
		return err
	}

	if err := log.SetupLog(); err != nil {
		return err
	}

	log.Infof("Current active profile: %s", Viper().GetString("env"))

	log.Infof("Current load config path: %s", getConfig().Gore.Path)

	if len(getConfig().Gore.Logger.Level) > 0 {
		log.SetLogLevel(getConfig().Gore.Logger.Level)
	} else {
		log.SetLogLevel(logrus.TraceLevel.String())
	}

	log.Infof("Current logger level: %s", log.GetLevel())

	if err := goreCache.Setup(getConfig().Gore.Cache); err != nil {
		return err
	}

	if err := goreEs.Setup(getConfig().Gore.Elasticsearch); err != nil {
		return err
	}

	if err := goreMongo.Setup(getConfig().Gore.Mongo); err != nil {
		return err
	}

	if err := goreMysql.Setup(getConfig().Gore.Mysql); err != nil {
		return err
	}

	if err := goreRedis.Setup(getConfig().Gore.Redis); err != nil {
		return err
	}

	if Viper().GetBool("gore.consul.enable") {
		if err := consul.Register(); err != nil {
			return err
		}
	}

	return nil
}

// Cmd return a root Command.
// preStartup is between gore.setup and server startup.
func Cmd(preStartup func(engine *gin.Engine) error) *cobra.Command {
	return cmd.New(func() error {
		if err := Setup(); err != nil {
			return err
		}
		if err := preStartup(goreGin.GetInstance()); err != nil {
			return err
		}
		return nil
	})
}

func InfraToken(c context.Context) (string, error) {
	return infratoken.Get(c, Viper().GetString("env"), Cache())
}

func Viper() *viper.Viper {
	return gonfig.GetViper()
}

func Gin() *gin.Engine {
	return goreGin.GetInstance()
}

func HttpClient() *http.Client {
	return goreHttp.GetInstance()
}

func Cache() *cache.Cache {
	return goreCache.GetInstance()
}

func CacheCustom(fn func(cfg *goreCache.Config) *cache.Cache) *cache.Cache {
	return fn(getConfig().Gore.Cache)
}

func Mongo() *mongo.Client {
	return goreMongo.GetInstance()
}

func MongoCustom(fn func(cfg *goreMongo.Config) *mongo.Client) *mongo.Client {
	return fn(getConfig().Gore.Mongo)
}

func Mysql() *sql.DB {
	return goreMysql.GetInstance()
}

func MysqlCustom(fn func(cfg *goreMysql.Config) *sql.DB) *sql.DB {
	return fn(getConfig().Gore.Mysql)
}

func Elasticsearch() *elastic.Client {
	return goreEs.GetInstance()
}

func ElasticsearchCustom(fn func(cfg *goreEs.Config) *elastic.Client) *elastic.Client {
	return fn(getConfig().Gore.Elasticsearch)
}

func KafkaStartConsumer(handler kafka.ConsumerMessageHandler) error {
	return goreKafka.StartConsumer(getConfig().Gore.Kafka, handler)
}

func KafkaStartConsumerCustom(fn func(cfg *goreKafka.Config) error) error {
	return fn(getConfig().Gore.Kafka)
}

func KafkaStartConsumers(handlers map[string]kafka.ConsumerMessageHandler) error {
	return goreKafka.SetupConsumers(getConfig().Gore.Kafka, handlers)
}

func KafkaStartConsumersCustom(fn func(cfg *goreKafka.Config) error) error {
	return fn(getConfig().Gore.Kafka)
}

func Redis() redis.UniversalClient {
	return goreRedis.GetInstance()
}

func RedisCustom(fn func(cfg *goreRedis.Config) redis.UniversalClient) redis.UniversalClient {
	return fn(getConfig().Gore.Redis)
}

func getConfig() *gonfig.Config {
	return gonfig.GetInstance()
}
