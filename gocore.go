package gore

import "git.tenvine.cn/backend/gore/log"

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

	if conf.Gore.Logger.Level != "" {
		log.SetLogLevel(conf.Gore.Logger.Level)
	}

	return nil
}
