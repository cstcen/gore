package mongo

import (
	"context"
	"git.tenvine.cn/backend/gore/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var (
	mgo *mongo.Client
)

type Config struct {
	Enable bool

	AppName  string `yaml:"app-name"`
	Username string
	Password string
	Hosts    []string
	Timeout  time.Duration
}

func GetInstance() *mongo.Client {
	return mgo
}

func Setup(cfg Config) error {
	if !cfg.Enable {
		return nil
	}

	var err error
	entry := log.WithField("host", cfg.Hosts)
	entry.Infof("Current mongo")

	clientOptions := newOptions(cfg)

	mgo, err = mongo.NewClient(clientOptions)
	if err != nil {
		entry.WithError(err).Warnf("Failed to create new mongodb")
		return err
	}

	if err = mgo.Connect(context.Background()); err != nil {
		entry.WithError(err).Warnf("Failed to connect mongodb")
		return err
	}

	return nil
}

func newOptions(c Config) *options.ClientOptions {
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
