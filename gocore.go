package gocore

import "github.com/sirupsen/logrus"

var (
	// 各环境对应的日志打印等级
	EnvLevelMap = map[Env]logrus.Level{
		EnvSDev0:  logrus.DebugLevel,
		EnvSDev:   logrus.DebugLevel,
		EnvDev:    logrus.DebugLevel,
		EnvDev2:   logrus.DebugLevel,
		EnvDev3:   logrus.DebugLevel,
		EnvIOS:    logrus.DebugLevel,
		EnvMod:    logrus.DebugLevel,
		EnvStg:    logrus.DebugLevel,
		EnvXingk5: logrus.DebugLevel,
		EnvXk5:    logrus.InfoLevel,
	}
)

// Setup 一键配置环境，日志和分解配置文件成struct
// env(required): 环境名称
// appName(required): 项目名称
// configOut(optional): 配置文件实例，configOut必须为指针，例如：new(conf.C)；
//                      为空，则不读取配置文件
func Setup(env string, appName string, configOut interface{}) error {

	err := SetupEnv(env)
	if err != nil {
		return err
	}

	err = SetupLog(appName)
	if err != nil {
		return err
	}

	if configOut != nil {
		err = UnmarshalConfig(configOut)
		if err != nil {
			return err
		}
	}

	return nil
}
