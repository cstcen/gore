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

    go get github.com/cstcen/gore

#### Config Example

    gore:
      logger:
        # 日志等级， trace, debug, info, warning, error, fatal, panic
        level: trace
      cache:
        # 是否开启缓存功能
        enable: false
        # 是否开启单点连接
        enableRing: false
        # 是否关闭状态检查
        disableStats: false
        # 实例名称
        appName: ${application}
        # host:port，如果开启单点连接模式，则请保持数组为一个元素
        hosts:
          - "localhost:6379"
        # 用户名
        username:
        # 密码
        password:
      elasticsearch:
        # 是否开启es功能
        enable: false
        config:
          # scheme://host:port
          url: http://localhost:9200
          # 用户名
          username:
          # 密码
          password:
      kafka:
        # 是否开启kafka功能
        enable: false
        # kafka版本
        version: 2.5.0
        # 负载均衡的策略, sticky, roundrobin, range
        assignor: range
        # 是否按最老排序
        oldest: false
        # 单个消费者时使用
        consumer:
          # host:port
          brokers:
            - "localhost:9092"
          # 主题
          topics:
            - ${application}
          # 分组
          group:
        # 多个消费者时使用，注意：`aaaa`要与代码中传入的处理消息方法的map-key对应，即map[string]kafka.ConsumerMessageHandler{"aaaa": ...}
        consumers:
          aaaa:
            # host:port
            brokers:
              - "localhost:9092"
            # 主题
            topics:
              - ${application}
            # 分组
            group:
      mongo:
        # 是否启用mongo功能
        enable: false
        # 实例名称
        appName: ${application}
        # 连接超时时间
        timeout: 30s
        # 用户名
        username:
        # 密码
        password:
        # host:port
        hosts:
          - "localhost:27017"
      mysql:
        # 是否开启mysql功能
        enable: false
        # 连接最大生命周期
        connMaxLifeTime: 3m
        # 最大打开的连接数
        maxOpenConns: 10
        # 最大空闲的连接数
        maxIdleConns: 10
        # DataSourceName
        Dsn:
          # `tcp` or `unix`
          protocol: tcp
          # 用户名
          username:
          # 密码
          password:
          # host:port
          address:
          # 数据库名称
          dbname:
          # 连接参数，例如：`?charset=UTF8&loc=UTC`
          params:
      redis:
        # 是否启用redis功能
        enable: false
        # 是否关闭集群模式
        disableCluster: false
        # host:port
        hosts:
          - "localhost:6379"
        # 用户名
        username:
        # 密码
        password:
      consul:
        enable: true
        host: https://i-consul-${profile}.xk5.com

#### Setup

    gore.Cmd(preStartup func(engine *gin.Engine) error) *cobra.Command

#### Infra Token

    gore.InfraToken(c context.Context) (string, error)

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
    