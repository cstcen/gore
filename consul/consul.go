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

	registration := api.AgentServiceRegistration{}
	if err := gonfig.Instance().UnmarshalKey("gore.consul.discovery", &registration); err != nil {
		return err
	}

	cli, err := api.NewClient(conf)
	if err != nil {
		return err
	}
	if err := cli.Agent().ServiceRegister(&registration); err != nil {
		return err
	}

	log.Infof("consul register service: %s", registration.Name)
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
