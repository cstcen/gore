package gore

import (
	"fmt"
	goreCache "git.tenvine.cn/backend/gore/db/cache"
	goreEs "git.tenvine.cn/backend/gore/db/es"
	goreKafka "git.tenvine.cn/backend/gore/db/kafka"
	goreMongo "git.tenvine.cn/backend/gore/db/mongo"
	goreMysql "git.tenvine.cn/backend/gore/db/mysql"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var conf = &Config{
	Gore: Gore{
		CommonPath:        "resource",
		CommonFileName:    "common.yml",
		CommonFileNameEnv: "common-%s.yml",
	},
}

var (
	ErrEnvEmpty       = errors.New("the env is empty")
	ErrConfigOutEmpty = errors.New("the config out is empty")

	// 当前文件夹路径
	curDir string
)

type Config struct {
	Gore Gore
}

type Gore struct {
	Path              string
	FileName          string `yaml:"file-name"`
	FileNameEnv       string `yaml:"file-name-env"`
	CommonPath        string `yaml:"common-path"`
	CommonFileName    string `yaml:"common-file-name"`
	CommonFileNameEnv string `yaml:"common-file-name-env"`
	// 日志配置
	Logger struct {
		// 日志打印等级
		Level string
	}
	Cache         goreCache.Config
	Elasticsearch goreEs.Config
	Mongo         goreMongo.Config
	Mysql         goreMysql.Config
	Kafka         goreKafka.Config
}

// SetupConfig 分解配置文件
// env: 环境名称
// out: 必须是指针
// config/config.yml
// config/config-{环境名称}.yml
func SetupConfig(env string, outPtr interface{}) error {

	if "" == env {
		return ErrEnvEmpty
	}

	if outPtr == nil {
		return ErrConfigOutEmpty
	}

	curDir = getCurDir()

	if err := unmarshalConfigCommon(); err != nil {
		return err
	}

	if err := unmarshalConfigCommonEnv(env); err != nil {
		return err
	}

	if err := unmarshalConfigDefault(); err != nil {
		return err
	}

	if err := unmarshalConfigCustom(outPtr); err != nil {
		return err
	}

	if env != "" {
		if err := unmarshalConfigCustomEnv(outPtr, env); err != nil {
			return err
		}
	}

	return nil
}

func GetConfig() *Config {
	return conf
}

func getCurDir() string {
	_, file, _, _ := runtime.Caller(0)
	return file[:strings.LastIndex(file, "/")]
}

func unmarshalConfigCommon() error {
	return unmarshal(filepath.Join(curDir, filepath.Join(conf.Gore.CommonPath, conf.Gore.CommonFileName)), conf)
}

func unmarshalConfigCommonEnv(env string) error {
	return unmarshal(filepath.Join(curDir, filepath.Join(conf.Gore.CommonPath, fmt.Sprintf(conf.Gore.CommonFileNameEnv, env))), conf)
}

func unmarshalConfigDefault() error {
	return unmarshal(filepath.Join(conf.Gore.Path, conf.Gore.FileName), conf)
}

func unmarshalConfigCustom(outPtr interface{}) error {
	return unmarshal(filepath.Join(conf.Gore.Path, conf.Gore.FileName), conf, outPtr)
}

func unmarshalConfigCustomEnv(outPtr interface{}, env string) error {
	return unmarshal(filepath.Join(conf.Gore.Path, fmt.Sprintf(conf.Gore.FileNameEnv, env)), conf, outPtr)
}

func unmarshal(filename string, outPtr ...interface{}) error {
	yml, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

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
