package gore

import (
	"database/sql"
	"git.tenvine.cn/backend/gore/db"
	goreCache "git.tenvine.cn/backend/gore/db/cache"
	goreEs "git.tenvine.cn/backend/gore/db/es"
	"git.tenvine.cn/backend/gore/db/kafka"
	goreKafka "git.tenvine.cn/backend/gore/db/kafka"
	goreMongo "git.tenvine.cn/backend/gore/db/mongo"
	goreMysql "git.tenvine.cn/backend/gore/db/mysql"
	goreRedis "git.tenvine.cn/backend/gore/db/redis"
	"git.tenvine.cn/backend/gore/log"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

// Setup 一键配置环境，日志和分解配置文件成struct
//
// env(required): 环境名称
// configOut(required): 配置文件实例，configOut必须为指针，例如：new(conf.C)
func Setup(env string, configOut interface{}) error {

	if err := SetupConfig(env, configOut); err != nil {
		return err
	}

	if err := log.SetupLog(); err != nil {
		return err
	}

	log.Infof("Current active profile: %s", env)

	log.Infof("Current load config path: %s", conf.Gore.Path)

	if len(conf.Gore.Logger.Level) > 0 {
		log.SetLogLevel(conf.Gore.Logger.Level)
	} else {
		log.SetLogLevel(logrus.TraceLevel.String())
	}

	log.Infof("Current logger level: %s", log.GetLevel())

	if err := goreCache.Setup(conf.Gore.Cache); err != nil {
		return err
	}

	if err := goreEs.Setup(conf.Gore.Elasticsearch); err != nil {
		return err
	}

	if err := goreMongo.Setup(conf.Gore.Mongo); err != nil {
		return err
	}

	if err := goreMysql.Setup(conf.Gore.Mysql); err != nil {
		return err
	}

	if err := goreRedis.Setup(conf.Gore.Redis); err != nil {
		return err
	}

	return nil
}

func CheckDB() *db.CheckResult {
	return db.Check(db.Config{
		Cache:         conf.Gore.Cache,
		Elasticsearch: conf.Gore.Elasticsearch,
		Mongo:         conf.Gore.Mongo,
		Mysql:         conf.Gore.Mysql,
		Redis:         conf.Gore.Redis,
	})
}

func Cache() *cache.Cache {
	return goreCache.GetInstance()
}

func CacheCustom(fn func(cfg goreCache.Config) *cache.Cache) *cache.Cache {
	return fn(conf.Gore.Cache)
}

func Mongo() *mongo.Client {
	return goreMongo.GetInstance()
}

func MongoCustom(fn func(cfg goreMongo.Config) *mongo.Client) *mongo.Client {
	return fn(conf.Gore.Mongo)
}

func Mysql() *sql.DB {
	return goreMysql.GetInstance()
}

func MysqlCustom(fn func(cfg goreMysql.Config) *sql.DB) *sql.DB {
	return fn(conf.Gore.Mysql)
}

func Elasticsearch() *elastic.Client {
	return goreEs.GetInstance()
}

func ElasticsearchCustom(fn func(cfg goreEs.Config) *elastic.Client) *elastic.Client {
	return fn(conf.Gore.Elasticsearch)
}

func KafkaStartConsumer(handler kafka.ConsumerMessageHandler) error {
	return goreKafka.StartConsumer(conf.Gore.Kafka, handler)
}

func KafkaStartConsumerCustom(fn func(cfg goreKafka.Config) error) error {
	return fn(conf.Gore.Kafka)
}

func KafkaStartConsumers(handlers map[string]kafka.ConsumerMessageHandler) error {
	return goreKafka.SetupConsumers(conf.Gore.Kafka, handlers)
}

func KafkaStartConsumersCustom(fn func(cfg goreKafka.Config) error) error {
	return fn(conf.Gore.Kafka)
}

func Redis() redis.UniversalClient {
	return goreRedis.GetInstance()
}

func RedisCustom(fn func(cfg goreRedis.Config) redis.UniversalClient) redis.UniversalClient {
	return fn(conf.Gore.Redis)
}
