## Gore

内部使用的基础依赖包

### Features

1. 提供一些基本的工具方法
2. 提供一些公用的常量
3. 提供InfraToken的获取方法
4. 提供规范的日志打印方法、日志保存、日志切割
5. 提供配置文件加载方法
6. 提供公用model：BaseResult,DataResult,PageResult
7. 提供gin初始化方法：var r *gin.Engine = gore.SetupGin(mode)，具体如下：
   1. 内置gin错误处理中间件：自动处理gin.Context.errorMsgs的错误，在业务层发现错误，可以通过gin.Context.Error(error)方法，把错误交给此中间件处理并响应
   2. 内置gin日志中间件：自动打印请求和响应的uri,header,body等信息
   3. 内置gin请求ID中间件：自动嵌入X-Request-ID到gin.Context.Keys
8. 提供日志X-Request-ID，想要打印带有X-Request-ID的业务日志，请使用(log.WithContext(c context.Context) *logrus.Entry)方法


### Usage

#### Install

    go get git.tenvine.cn/backend/gore
    
#### Config Example

    gore:
      path: config
      file-name: config.yml
      file-name-env: config-%s.yml
      common-path: config
      common-file-name: common.yml
      common-file-name-env: common-%s.yml
      logger:
        level: trace
      cache:
        enable: false
        enable-ring: false
        app-name: gore
        hosts: 10.251.43.21:6379
        username:
        password: sdev0#redis#3aPVNy
      elasticsearch:
        enable: false
        url: http://10.251.104.6:9200
        index:
        username:
        shards: 1
        replicas: 0
        password: allDev#es#3aPVNy
      mongo:
        enable: false
        app-name: gore
        username:
        password:
        hosts: 10.251.104.15:27017
        timeout: 30s
      mysql:
        enable: false
        conn-max-life-time: 3m
        max-open-conns: 10
        max-idle-conns: 10
        dsn:
          username:
          password:
          protocol: tcp
          address: localhost:4725
          dbname:
          params: ?charset=UTF8&loc=UTC
      kafka:
        enable: false
        version: 2.5.0
        assignor: range
        oldest: false
        consumer:
          brokers:
          topics:
          group:
        consumers:
          member:
            brokers:
            topics:
            group:
    
#### Setup

    Setup(env string, configOut interface{}) error

#### Infra Token

    // Setup 一键配置环境，日志和分解配置文件成struct
    //
    // env(required): 环境名称
    // configOut(required): 配置文件实例，configOut必须为指针，例如：new(conf.C)
    gore.GetInfraToken() (*gore.InfraTokenResponse, error)
    
#### Config context

    gore:
        path: ./config
        fileName: config.yml
        fileNameEnv: sdev0
        logger:
            level: trace

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
    log.Trace("test")
    log.Tracef("test %s", "Tracef")
    
    log.ErrorE(err error)
    log.WarnE(err error)
    