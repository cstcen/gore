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
	"git.tenvine.cn/backend/gore/usertoken"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/olivere/elastic"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func Setup() error {

	if err := gonfig.Setup(); err != nil {
		return err
	}

	if err := log.Setup(); err != nil {
		return err
	}

	log.Infof("Current active profile: %s", Viper().GetString("env"))

	log.Infof("Current load config path: %s", Viper().GetString("gore.path"))

	log.Infof("Current logger level: %s", log.GetLevel())

	if err := goreGin.Setup(); err != nil {
		return err
	}

	if err := goreCache.Setup(); err != nil {
		return err
	}

	if err := goreEs.Setup(); err != nil {
		return err
	}

	if err := goreMongo.Setup(); err != nil {
		return err
	}

	if err := goreMysql.Setup(); err != nil {
		return err
	}

	if err := goreRedis.Setup(); err != nil {
		return err
	}

	if err := consul.Setup(); err != nil {
		return err
	}

	if consul.Enable() {
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
	return infratoken.Get(c)
}

func UserTokenVerification(token string) (*usertoken.Member, error) {
	return usertoken.Check(token)
}

func Viper() *viper.Viper {
	return gonfig.Instance()
}

func Gin() *gin.Engine {
	return goreGin.GetInstance()
}

func HttpClient() *http.Client {
	return goreHttp.GetInstance()
}

func HttpInternalPost(c context.Context, url, contentType string, body interface{}, expectedPtr interface{}) error {
	return goreHttp.InternalPost(c, url, contentType, body, expectedPtr, InfraToken)
}

func HttpPost(url, contentType string, body interface{}, expectedPtr interface{}) error {
	return goreHttp.Post(url, contentType, body, expectedPtr)
}

func HttpInternalGet(c context.Context, url string, expectedPtr interface{}) error {
	return goreHttp.InternalGet(c, url, expectedPtr, InfraToken)
}

func HttpGet(url string, expectedPtr interface{}) error {
	return goreHttp.Get(url, expectedPtr)
}

func HttpHead(url string, expectedPtr interface{}) error {
	return goreHttp.Get(url, expectedPtr)
}

func Cache() *cache.Cache {
	return goreCache.Instance()
}

func CacheCustom(setup func() *cache.Cache) *cache.Cache {
	return setup()
}

func Mongo() *mongo.Client {
	return goreMongo.Instance()
}

func MongoCustom(setup func() *mongo.Client) *mongo.Client {
	return setup()
}

func Mysql() *sql.DB {
	return goreMysql.Instance()
}

func MysqlCustom(setup func() *sql.DB) *sql.DB {
	return setup()
}

func Elasticsearch() *elastic.Client {
	return goreEs.Instance()
}

func ElasticsearchCustom(setup func() *elastic.Client) *elastic.Client {
	return setup()
}

func KafkaStartConsumer(handler kafka.ConsumerMessageHandler) error {
	return goreKafka.StartConsumer(handler)
}

func KafkaStartConsumerCustom(setup func() error) error {
	return setup()
}

func KafkaStartConsumers(handlers map[string]kafka.ConsumerMessageHandler) error {
	return goreKafka.SetupConsumers(handlers)
}

func KafkaStartConsumersCustom(setup func() error) error {
	return setup()
}

func Redis() redis.UniversalClient {
	return goreRedis.Instance()
}

func RedisCustom(setup func() redis.UniversalClient) redis.UniversalClient {
	return setup()
}
