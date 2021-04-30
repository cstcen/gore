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
func Setup(env string, appName string, configStruct interface{}) error {

	err := SetupEnv(env)
	if err != nil {
		return err
	}

	err = SetupLog(appName)
	if err != nil {
		return err
	}

	err = UnmarshalConfig(configStruct)
	if err != nil {
		return err
	}

	return nil
}
