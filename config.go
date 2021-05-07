package gore

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

var conf = &Config{
	Gore: Gore{
		Path:        "config/",
		FileName:    "config.yml",
		FileNameEnv: "config-%s.yml",
	},
}

var (
	ErrEnvEmpty       = errors.New("the env is empty")
	ErrConfigOutEmpty = errors.New("the config out is empty")
)

type Config struct {
	Gore Gore
}

type Gore struct {
	Path        string
	FileName    string
	FileNameEnv string
	Logger      struct {
		Level string
	}
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

func unmarshalConfigCustomEnv(value interface{}, env string) error {
	return unmarshal(fmt.Sprintf(conf.Gore.Path+conf.Gore.FileNameEnv, env), value)
}

func unmarshalConfigDefault() error {
	return unmarshal(conf.Gore.Path+conf.Gore.FileName, conf)
}

func unmarshalConfigCustom(outPtr interface{}) error {
	return unmarshal(conf.Gore.Path+conf.Gore.FileName, outPtr)
}

func unmarshal(filename string, outPtr interface{}) error {
	yml, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yml, outPtr)
	if err != nil {
		return err
	}

	return nil
}
