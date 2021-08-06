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
	crypt "github.com/bketelsen/crypt/config"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"strings"
)

var (
	ErrEnvEmpty = errors.New("the env is empty")

	vp = viper.New()

	conf *Config

	storeKeyFns = []func(env, appName string) string{
		func(env, appName string) string {
			return fmt.Sprintf("config/go/%s/data", appName)
		},
		func(env, appName string) string {
			return fmt.Sprintf("config/go/%s,%s/data", appName, env)
		},
	}
)

// Setup 读取顺序为：
// 假如env=dev, appName=gore
// 先读取consul K/V store
// 1. config/application/data
// 2. config/application,dev/data
// 3. config/gore/data
// 4. config/gore,dev/data
// 再读取项目内配置文件
// 1. config/config.yml
// 2. config/config-dev.yml
// 越靠后，配置优先级越高
func Setup() error {

	env := vp.GetString("env")
	if len(env) == 0 {
		return ErrEnvEmpty
	}
	appName := vp.GetString("name")

	endpoint := "10.251.110.122:8500"

	vp.SetConfigType("yaml")

	name := "application"
	if err := readRemoteConfig(env, name, endpoint); err != nil {
		return err
	}
	if len(appName) > 0 {
		if err := readRemoteConfig(env, appName, endpoint); err != nil {
			return err
		}
	}

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

	for _, key := range vp.AllKeys() {
		val, ok := vp.Get(key).(string)
		if !ok {
			continue
		}
		val = strings.ReplaceAll(val, "${profile}", env)
		val = strings.ReplaceAll(val, "${application}", appName)
		vp.Set(key, val)
	}

	if err := vp.UnmarshalKey("gore", conf.Gore); err != nil {
		return err
	}

	return nil
}

func readRemoteConfig(env string, appName string, endpoint string) error {
	var cm crypt.ConfigManager
	var err error

	cm, err = crypt.NewStandardConsulConfigManager([]string{endpoint})
	if err != nil {
		return err
	}

	for _, fn := range storeKeyFns {
		path := fn(env, appName)
		b, err := cm.Get(path)
		if err != nil {
			continue
		}
		if err := vp.MergeConfig(bytes.NewReader(b)); err != nil {
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
			//Logger:      &log.Config{Level: "trace"},
			//Cache: &goreCache.Config{
			//	Enable:       false,
			//	EnableRing:   false,
			//	DisableStats: false,
			//	AppName:      "gore",
			//	Hosts:        []string{"localhost:6379"},
			//	Username:     "",
			//	Password:     "",
			//},
			//Elasticsearch: &goreEs.Config{
			//	Enable: false,
			//	Config: esConfig.Config{
			//		Index: "",
			//	},
			//},
			//Kafka: &goreKafka.Config{
			//	Enable:   false,
			//	Version:  "2.5.0",
			//	Assignor: "range",
			//	Oldest:   false,
			//},
			//Mongo: &goreMongo.Config{
			//	Enable:  false,
			//	AppName: "gore",
			//	Timeout: 30 * time.Second,
			//},
			//Mysql: &goreMysql.Config{
			//	Enable:          false,
			//	ConnMaxLifeTime: 3 * time.Minute,
			//	MaxOpenConns:    10,
			//	MaxIdleConns:    10,
			//	Dsn: goreMysql.DataSourceName{
			//		Protocol: "tcp",
			//		Params:   "?charset=UTF8&loc=UTC",
			//	},
			//},
			//Redis: &goreRedis.Config{
			//	Enable:         false,
			//	DisableCluster: false,
			//},
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
