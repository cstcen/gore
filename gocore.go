package gore

import (
	"git.tenvine.cn/backend/gore/log"
	"github.com/sirupsen/logrus"
)

// Setup 一键配置环境，日志和分解配置文件成struct
//
// env(required): 环境名称
// configOut(required): 配置文件实例，configOut必须为指针，例如：new(conf.C)
func Setup(env string, configOut interface{}) error {

	if err := SetupConfig(env, configOut); err != nil {
		return err
	}

	if err := log.SetupLog(); err != nil {
		return err
	}

	log.Infof("Current active profile: %s", env)

	log.Infof("Current load config path: %s", conf.Gore.Path)

	if conf.Gore.Logger.Level != "" {
		log.SetLogLevel(conf.Gore.Logger.Level)
	} else {
		log.SetLogLevel(logrus.TraceLevel.String())
	}

	log.Infof("Current logger level: %s", log.GetLevel())

	return nil
}
