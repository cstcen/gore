package gore

import (
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
// config/config.yml
// config/config-{环境名称}.yml
func SetupConfig(env string, out interface{}) error {

	if "" == env {
		return ErrEnvEmpty
	}

	if out == nil {
		return ErrConfigOutEmpty
	}

	if err := unmarshalConfigDefault(); err != nil {
		return err
	}

	if err := unmarshalConfigCustom(out); err != nil {
		return err
	}

	if env != "" {
		if err := unmarshalConfigCustomEnv(out, env); err != nil {
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

func unmarshalConfigCustom(value interface{}) error {
	return unmarshal(conf.Gore.Path+conf.Gore.FileName, value)
}

func unmarshal(filename string, value interface{}) error {
	yml, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yml, value)
	if err != nil {
		return err
	}
	return nil
}
