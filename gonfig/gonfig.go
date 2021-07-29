package gonfig

import (
	"bytes"
	"fmt"
	goreCache "git.tenvine.cn/backend/gore/db/cache"
	goreEs "git.tenvine.cn/backend/gore/db/es"
	goreKafka "git.tenvine.cn/backend/gore/db/kafka"
	goreMongo "git.tenvine.cn/backend/gore/db/mongo"
	goreMysql "git.tenvine.cn/backend/gore/db/mysql"
	goreRedis "git.tenvine.cn/backend/gore/db/redis"
	"git.tenvine.cn/backend/gore/log"
	esConfig "github.com/olivere/elastic/config"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"time"
)

var (
	ErrEnvEmpty = errors.New("the env is empty")

	vp = viper.New()

	conf *Config
)

func Setup(env, appName string) error {

	if len(env) == 0 {
		return ErrEnvEmpty
	}

	Initialize(env, appName)

	vp.SetConfigType("yaml")

	if err := unmarshalConfigDefault(); err != nil {
		return err
	}

	if err := unmarshalConfigCustom(); err != nil {
		return err
	}

	if len(env) > 0 {
		if err := unmarshalConfigCustomEnv(env); err != nil {
			return err
		}
	}

	if err := vp.UnmarshalKey("gore", conf.Gore); err != nil {
		return err
	}

	return nil
}

type Config struct {
	Gore *Gore
}

type Gore struct {
	Path        string
	Filename    string
	FilenameEnv string
	// 日志配置
	Logger        *log.Config
	Cache         *goreCache.Config
	Elasticsearch *goreEs.Config
	Kafka         *goreKafka.Config
	Mongo         *goreMongo.Config
	Mysql         *goreMysql.Config
	Redis         *goreRedis.Config
}

func GetInstance() *Config {
	return conf
}

func GetViper() *viper.Viper {
	return vp
}

func init() {
	conf = &Config{
		Gore: &Gore{
			Path:        "config",
			Filename:    "config.yml",
			FilenameEnv: "config-%s.yml",
			Logger:      &log.Config{Level: "trace"},
			Cache: &goreCache.Config{
				Enable:       false,
				EnableRing:   false,
				DisableStats: false,
				AppName:      "gore",
			},
			Elasticsearch: &goreEs.Config{
				Enable: false,
				Config: esConfig.Config{
					Index: "",
				},
			},
			Kafka: &goreKafka.Config{
				Enable:   false,
				Version:  "2.5.0",
				Assignor: "range",
				Oldest:   false,
			},
			Mongo: &goreMongo.Config{
				Enable:  false,
				AppName: "gore",
				Timeout: 30 * time.Second,
			},
			Mysql: &goreMysql.Config{
				Enable:          false,
				ConnMaxLifeTime: 3 * time.Minute,
				MaxOpenConns:    10,
				MaxIdleConns:    10,
				Dsn: goreMysql.DataSourceName{
					Protocol: "tcp",
					Params:   "?charset=UTF8&loc=UTC",
				},
			},
			Redis: &goreRedis.Config{
				Enable:         false,
				DisableCluster: false,
			},
		},
	}
}

func unmarshalConfigDefault() error {
	out, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}

	if err := vp.ReadConfig(bytes.NewBuffer(out)); err != nil {
		return err
	}
	return nil
}

func unmarshalConfigCustom() error {
	return unmarshal(filepath.Join(conf.Gore.Path, conf.Gore.Filename))
}

func unmarshalConfigCustomEnv(env string) error {
	return unmarshal(filepath.Join(conf.Gore.Path, fmt.Sprintf(conf.Gore.FilenameEnv, env)))
}

func unmarshal(filename string) error {
	yml, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	if err := vp.MergeConfig(bytes.NewBuffer(yml)); err != nil {
		return err
	}

	return nil
}
