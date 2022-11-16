package gore

import (
	"context"
	"database/sql"
	"git.tenvine.cn/backend/gore/auth"
	"git.tenvine.cn/backend/gore/cmd"
	"git.tenvine.cn/backend/gore/command"
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
	"git.tenvine.cn/backend/gore/middleware"
	"git.tenvine.cn/backend/gore/util"
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

	if err := SetupBase(); err != nil {
		return err
	}

	if err := goreHttp.Setup(); err != nil {
		return err
	}

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

	if err := consul.Register(); err != nil {
		return err
	}

	return nil
}

func SetupBase() error {
	if err := gonfig.Setup(); err != nil {
		return err
	}

	if err := log.Setup(); err != nil {
		return err
	}

	log.Infof("Current active profile: %s", Viper().GetString("env"))

	log.Infof("Current load config path: %s", Viper().GetString("gore.path"))

	return nil
}

func SetupGin() error {
	return goreGin.Setup()
}

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

func RootCmd() *cobra.Command {
	return cmd.GetInstance()
}

func Gin() *gin.Engine {
	return goreGin.GetInstance()
}

func InfraToken(c context.Context) (string, error) {
	return infratoken.Get(c)
}

func UserTokenVerification(token string) (*auth.Member, error) {
	return auth.ExternalCheck(context.WithValue(context.Background(), util.RequestIDContextKey, util.GenerateRequestID()), token)
}

func TokenVerification(token string) (*auth.Member, error) {
	return auth.InternalCheck(context.WithValue(context.Background(), util.RequestIDContextKey, util.GenerateRequestID()), token)
}

func ExternalTokenVerification(token string) (*auth.Member, error) {
	return auth.ExternalCheck(context.WithValue(context.Background(), util.RequestIDContextKey, util.GenerateRequestID()), token)
}

func InternalTokenVerification(token string) (*auth.Member, error) {
	return auth.InternalCheck(context.WithValue(context.Background(), util.RequestIDContextKey, util.GenerateRequestID()), token)
}

func Viper() *viper.Viper {
	return gonfig.Instance()
}

func HttpClient() *http.Client {
	return goreHttp.GetInstance()
}

func HttpInternalPost(c context.Context, url, contentType string, body any, expectedPtr any) error {
	return goreHttp.InternalPost(c, url, contentType, body, expectedPtr, InfraToken)
}

func HttpPost(c context.Context, url, contentType string, body any, expectedPtr any) error {
	return goreHttp.Post(c, url, contentType, body, expectedPtr)
}

func HttpInternalGet(c context.Context, url string, expectedPtr any) error {
	return goreHttp.InternalGet(c, url, expectedPtr, InfraToken)
}

func HttpGet(c context.Context, url string, expectedPtr any) error {
	return goreHttp.Get(c, url, expectedPtr)
}

func HttpHead(c context.Context, url string, expectedPtr any) error {
	return goreHttp.Head(c, url, expectedPtr)
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

func MongoDatabase() *mongo.Database {
	return goreMongo.Database()
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

func KafkaStartConsumers(handlers map[string]kafka.ConsumerMessageHandler) error {
	return goreKafka.StartupConsumers(handlers)
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

func MiddlewareRequestID(handler http.Handler) http.Handler {
	return middleware.SetupRequestID(handler)
}
func MiddlewareTrace(handler http.Handler, skipLogResp func(path string) bool) http.Handler {
	return middleware.SetupTrace(handler, skipLogResp)
}
func MiddlewareRecovery(handler http.Handler) http.Handler {
	return middleware.SetupRecovery(handler)
}

func MiddlewareDB(handler http.Handler) http.Handler {
	return middleware.SetupDB(handler)
}

func SetupGorm() error {
	return goreMysql.SetupGorm()
}

func SetupHttp() error {
	return goreHttp.Setup()
}

func SetDefaultConfig(opt *command.Args) {
	Viper().Set("name", opt.Name)
	Viper().Set("env", opt.Env)
	Viper().Set("consul", opt.Consul)
	Viper().Set("log", opt.Log)
}

func Debugf(format string, v ...any) {
	log.Debugf(format, v...)
}
