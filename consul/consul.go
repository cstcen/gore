package consul

import (
	"git.tenvine.cn/backend/gore/gonfig"
	goreHttp "git.tenvine.cn/backend/gore/http"
	"git.tenvine.cn/backend/gore/log"
	"github.com/hashicorp/consul/api"
)

var (
	conf *api.Config
)

func Setup() error {
	url := gonfig.Instance().GetString("gore.consul.discovery.host")
	conf = &api.Config{
		Address:    url,
		HttpClient: goreHttp.GetInstance(),
	}
	return nil
}

func Register() error {

	cli, err := api.NewClient(conf)
	if err != nil {
		return err
	}

	appName := gonfig.Instance().GetString("name")
	env := gonfig.Instance().GetString("env")
	addr := gonfig.Instance().GetString("SERVER_ID")
	registration := &api.AgentServiceRegistration{Name: appName, Tags: []string{env}, Address: addr}
	if err := cli.Agent().ServiceRegister(registration); err != nil {
		return err
	}

	log.Infof("consul register service: %s", appName)
	return nil

}

func Deregister() error {

	cli, err := api.NewClient(conf)
	if err != nil {
		return err
	}

	appName := gonfig.Instance().GetString("name")
	if err := cli.Agent().ServiceDeregister(appName); err != nil {
		return err
	}

	log.Infof("consul deregister service: %s", appName)
	return nil
}

func Enable() bool {
	return gonfig.Instance().GetBool("gore.consul.enable")
}
