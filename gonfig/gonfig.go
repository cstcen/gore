package gonfig

import (
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
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"time"
)

var (
	ErrEnvEmpty       = errors.New("the env is empty")
	ErrConfigOutEmpty = errors.New("the config out is empty")

	conf    *Config
	confMap = make(map[string]interface{})
)

func Setup(env string, outPtr ...interface{}) error {

	if "" == env {
		return ErrEnvEmpty
	}

	if outPtr == nil {
		return ErrConfigOutEmpty
	}

	Initialize(env)

	if err := unmarshalConfigDefault(); err != nil {
		return err
	}

	if err := unmarshalConfigCustom(outPtr...); err != nil {
		return err
	}

	if len(env) > 0 {
		if err := unmarshalConfigCustomEnv(env, outPtr...); err != nil {
			return err
		}
	}

	return nil
}

type Config struct {
	Gore *Gore
}

type Gore struct {
	Path        string
	FileName    string `yaml:"file-name"`
	FileNameEnv string `yaml:"file-name-env"`
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

func GetInstanceMap(key string) (interface{}, bool) {
	v, ok := confMap[key]
	return v, ok
}

func init() {
	conf = &Config{
		Gore: &Gore{
			Path:        "config",
			FileName:    "config.yml",
			FileNameEnv: "config-%s.yml",
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
	return unmarshal(filepath.Join(conf.Gore.Path, conf.Gore.FileName))
}

func unmarshalConfigCustom(outPtr ...interface{}) error {
	return unmarshal(filepath.Join(conf.Gore.Path, conf.Gore.FileName), outPtr...)
}

func unmarshalConfigCustomEnv(env string, outPtr ...interface{}) error {
	return unmarshal(filepath.Join(conf.Gore.Path, fmt.Sprintf(conf.Gore.FileNameEnv, env)), outPtr...)
}

func unmarshal(filename string, outPtr ...interface{}) error {
	yml, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	outPtr = append(outPtr, conf, confMap)

	if len(outPtr) > 0 {
		for i, p := range outPtr {
			err = yaml.Unmarshal(yml, p)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("gore %s: out ptr [#%2d]", filename, i))
			}
		}
	}

	return nil
}
