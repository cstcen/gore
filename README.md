## Go Core
内部使用的基础依赖包
### Features
1. 提供一些基本的工具方法
2. 存放一些公用的常量
3. 提供InfraToken的获取方法
4. 提供规范的日志打印方法、日志保存、日志切割


### Usage

#### Install
`go get git.tenvine.cn/backend/gocore`

#### Infra Token

`gocore.GetInfraToken(环境名称) (*gocore.InfraTokenResponse, error)`

