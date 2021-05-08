## Gore

内部使用的基础依赖包

### Features

1. 提供一些基本的工具方法
2. 存放一些公用的常量
3. 提供InfraToken的获取方法
4. 提供规范的日志打印方法、日志保存、日志切割
5. 提供配置文件加载方法


### Usage

#### Install

    go get git.tenvine.cn/backend/gore

#### Infra Token

    gore.GetInfraToken() (*gore.InfraTokenResponse, error)

#### Loading Config File

在程序执行前，把对应名称的配置文件存放到路径：./config/

在程序执行时，自动加载以下文件

    ./config/config.yml 和 ./config/config-{环境名称}.yml
    


需要传两个参数，环境名称和配置文件输出对象的**指针**

    gore.SetupConfig(env string, outPtr interface{}) error
    
    t := new(Tes)
	err := gore.SetupConfig("sdev0", t)

#### Setup Logger

假设，程序build之后的执行文件路径为：/xk5/app/gdis/

那么，日志存放的默认路径便是：'/xk5/logs/' + 'gdis' + '/gdis.log'

所以，'/xk5/logs/'是**固定值**，'/gdis.log'也是**固定值**，中间的'gdis'取的是执行文件的所在文件夹**名称**

    err := log.SetupLog()
    
    log.Info("test")
    log.Infof("test %s", "Infof")
    log.Warn("test")
    log.Warnf("test %s", "Warnf")
    log.Debug("test")
    log.Debugf("test %s", "Debugf")
    log.Error("test")
    log.Errorf("test %s", "Errorf")
    
    log.ErrorE(err error)
    log.WarnE(err error)
    
