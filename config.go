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
func UnmarshalConfig(value interface{}) error {

	err := unmarshalFile(value)
	if err != nil {
		return err
	}

	err = unmarshalFileEnv(value)
	if err != nil {
		return err
	}

	return nil
}

func unmarshalFileEnv(value interface{}) error {
	yml, err := os.ReadFile(fmt.Sprintf(ConfigPath+configFileNameEnv, envCurrentName))
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yml, value)
	if err != nil {
		return err
	}
	return nil
}

func unmarshalFile(value interface{}) error {
	yml, err := os.ReadFile(ConfigPath + configFileName)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yml, value)
	if err != nil {
		return err
	}
	return nil
}
