package gocore

// Setup 一键配置环境，日志和分解配置文件成struct
// env(required): 环境名称
// appName(required): 项目名称
// configOut(optional): 配置文件实例，configOut必须为指针，例如：new(conf.C)；
//                      为空，则不读取配置文件
func Setup(env string, appName string, configOut interface{}) error {

	err := SetupLog(env, appName)
	if err != nil {
		return err
	}

	if configOut != nil {
		err = UnmarshalConfig(env, configOut)
		if err != nil {
			return err
		}
	}

	return nil
}
