package gocore

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

var (
	ConfigPath = "config/"

	configFileName    = "config.yml"
	configFileNameEnv = "config-%s.yml"
)

// UnmarshalConfig 分解配置文件
// config/config.yml
// config/config-{环境名称}.yml
func UnmarshalConfig(env string, value interface{}) error {

	err := unmarshalFile(value)
	if err != nil {
		return err
	}

	err = unmarshalFileEnv(value, env)
	if err != nil {
		return err
	}

	return nil
}

func unmarshalFileEnv(value interface{}, env string) error {
	return unmarshal(fmt.Sprintf(ConfigPath+configFileNameEnv, env), value)
}

func unmarshalFile(value interface{}) error {
	return unmarshal(ConfigPath+configFileName, value)
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
