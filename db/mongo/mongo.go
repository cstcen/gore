package mongo

import (
	"context"
	"git.tenvine.cn/backend/gore/gonfig"
	"git.tenvine.cn/backend/gore/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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

func NewConfig() *Config {
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

func Setup() error {
	cfg := NewConfig()

	if !cfg.Enable {
		return nil
	}

	var err error
	entry := log.WithField("host", cfg.Hosts)
	entry.Infof("Current mongo dbname: %s", cfg.Dbname)

	clientOptions := newOptions(cfg)

	mgo, err = mongo.NewClient(clientOptions)
	if err != nil {
		entry.WithError(err).Warnf("Failed to create new mongodb")
		return err
	}

	if err = mgo.Connect(context.Background()); err != nil {
		entry.WithError(err).Warnf("Failed to connect mongodb [$s|$s]", cfg.Username, cfg.Password)
		return err
	}

	if err = mgo.Ping(context.Background(), readpref.Primary()); err != nil {
		entry.WithError(err).Warnf("Failed to ping mongodb [%s|%s]", cfg.Username, cfg.Password)
		return err
	}

	db = mgo.Database(cfg.Dbname)

	return nil
}

func newOptions(c *Config) *options.ClientOptions {
	clientOptions := options.Client()
	clientOptions.SetAppName(c.AppName)
	clientOptions.SetAuth(options.Credential{
		AuthSource:  c.AppName,
		Username:    c.Username,
		Password:    c.Password,
		PasswordSet: true,
	})
	clientOptions.SetHosts(c.Hosts)
	clientOptions.SetConnectTimeout(c.Timeout)
	return clientOptions
}
