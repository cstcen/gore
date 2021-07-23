package gore

import (
	"database/sql"
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
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

// Setup 一键配置环境，日志和分解配置文件成struct
//
// env(required): 环境名称
// ptrOut(required): 配置文件实例，ptrOut必须为指针，例如：new(GetConfig().C)
func Setup(env string, ptrOut ...interface{}) error {

	if err := gonfig.Setup(env, ptrOut...); err != nil {
		return err
	}

	if err := log.SetupLog(); err != nil {
		return err
	}

	log.Infof("Current active profile: %s", env)

	log.Infof("Current load config path: %s", GetConfig().Gore.Path)

	if len(GetConfig().Gore.Logger.Level) > 0 {
		log.SetLogLevel(GetConfig().Gore.Logger.Level)
	} else {
		log.SetLogLevel(logrus.TraceLevel.String())
	}

	log.Infof("Current logger level: %s", log.GetLevel())

	if err := goreCache.Setup(GetConfig().Gore.Cache); err != nil {
		return err
	}

	if err := goreEs.Setup(GetConfig().Gore.Elasticsearch); err != nil {
		return err
	}

	if err := goreMongo.Setup(GetConfig().Gore.Mongo); err != nil {
		return err
	}

	if err := goreMysql.Setup(GetConfig().Gore.Mysql); err != nil {
		return err
	}

	if err := goreRedis.Setup(GetConfig().Gore.Redis); err != nil {
		return err
	}

	return nil
}

func GetConfig() *gonfig.Config {
	return gonfig.GetInstance()
}

func GetConfigValue(key string) (interface{}, bool) {
	return gonfig.GetInstanceMap(key)
}

func GetInfraToken(env string) (*infratoken.Response, error) {
	return infratoken.GetInstance(env)
}

func Gin(ginMode string) *gin.Engine {
	return goreGin.Setup(ginMode)
}

func HttpClient() *http.Client {
	return goreHttp.GetInstance()
}

func Cache() *cache.Cache {
	return goreCache.GetInstance()
}

func CacheCustom(fn func(cfg *goreCache.Config) *cache.Cache) *cache.Cache {
	return fn(GetConfig().Gore.Cache)
}

func Mongo() *mongo.Client {
	return goreMongo.GetInstance()
}

func MongoCustom(fn func(cfg *goreMongo.Config) *mongo.Client) *mongo.Client {
	return fn(GetConfig().Gore.Mongo)
}

func Mysql() *sql.DB {
	return goreMysql.GetInstance()
}

func MysqlCustom(fn func(cfg *goreMysql.Config) *sql.DB) *sql.DB {
	return fn(GetConfig().Gore.Mysql)
}

func Elasticsearch() *elastic.Client {
	return goreEs.GetInstance()
}

func ElasticsearchCustom(fn func(cfg *goreEs.Config) *elastic.Client) *elastic.Client {
	return fn(GetConfig().Gore.Elasticsearch)
}

func KafkaStartConsumer(handler kafka.ConsumerMessageHandler) error {
	return goreKafka.StartConsumer(GetConfig().Gore.Kafka, handler)
}

func KafkaStartConsumerCustom(fn func(cfg *goreKafka.Config) error) error {
	return fn(GetConfig().Gore.Kafka)
}

func KafkaStartConsumers(handlers map[string]kafka.ConsumerMessageHandler) error {
	return goreKafka.SetupConsumers(GetConfig().Gore.Kafka, handlers)
}

func KafkaStartConsumersCustom(fn func(cfg *goreKafka.Config) error) error {
	return fn(GetConfig().Gore.Kafka)
}

func Redis() redis.UniversalClient {
	return goreRedis.GetInstance()
}

func RedisCustom(fn func(cfg *goreRedis.Config) redis.UniversalClient) redis.UniversalClient {
	return fn(GetConfig().Gore.Redis)
}
