package mongo

import (
	"context"
	"fmt"
	"github.com/cstcen/gore/gonfig"
	"github.com/cstcen/gore/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"strings"
	"time"
)

var (
	mgo *mongo.Client
	db  *mongo.Database
)

type Config struct {
	Enable bool

	AppName  string
	Username string
	Password string
	Hosts    []string
	Dbname   string
	Timeout  time.Duration
}

func Instance() *mongo.Client {
	return mgo
}

func Database() *mongo.Database {
	return db
}

func DefaultConfig() *Config {
	viper := gonfig.Instance()
	cfg := &Config{
		Enable:   viper.GetBool("gore.mongo.enable"),
		AppName:  viper.GetString("gore.mongo.appName"),
		Username: viper.GetString("gore.mongo.username"),
		Password: viper.GetString("gore.mongo.password"),
		Hosts:    viper.GetStringSlice("gore.mongo.hosts"),
		Dbname:   viper.GetString("gore.mongo.dbname"),
		Timeout:  viper.GetDuration("gore.mongo.timeout"),
	}
	return cfg
}

func SetupDefault() error {
	cfg := DefaultConfig()
	return Setup(cfg)
}

func Setup(cfg *Config) error {

	if !cfg.Enable {
		return nil
	}

	var err error
	logPrefix := fmt.Sprintf("host: %v", cfg.Hosts)
	log.Infof("[%s] Current mongo dbname: %s", logPrefix, cfg.Dbname)

	clientOptions := newOptions(cfg)

	mgo, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Warningf("[%s] failed to connect mongodb [%s|%s]", logPrefix, cfg.Username, cfg.Password)
		return err
	}

	if err = mgo.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Warningf("[%s] failed to ping mongodb [%s|%s]", logPrefix, cfg.Username, cfg.Password)
		return err
	}

	db = mgo.Database(cfg.Dbname)

	return nil
}

func Shutdown() error {
	return mgo.Disconnect(context.TODO())
}

func newOptions(c *Config) *options.ClientOptions {
	clientOptions := options.Client()
	credential := options.Credential{
		AuthSource:  c.Dbname,
		Username:    c.Username,
		Password:    c.Password,
		PasswordSet: true,
	}
	loggerOptions := options.Logger().SetComponentLevel(options.LogComponentCommand, options.LogLevelDebug).SetSink(&Sink{})
	clientOptions.ApplyURI(strings.Join(c.Hosts, ",")).
		SetAuth(credential).
		SetConnectTimeout(c.Timeout).
		SetBSONOptions(&options.BSONOptions{
			UseJSONStructTags: true,
			NilMapAsEmpty:     true,
			NilSliceAsEmpty:   true,
		}).
		SetLoggerOptions(loggerOptions)
	return clientOptions
}

type Sink struct {
}

func (s *Sink) Info(level int, message string, keysAndValues ...interface{}) {
	log.Infof(message, keysAndValues...)
}

func (s *Sink) Error(err error, message string, keysAndValues ...interface{}) {
	log.Errorf(fmt.Sprintf("err: %s, message: %s", err, message), keysAndValues...)
}
