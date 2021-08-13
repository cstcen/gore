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

func init() {
	url := gonfig.GetViper().GetString("gore.consul.host")
	conf = &api.Config{
		Address:    url,
		Scheme:     "https",
		HttpClient: goreHttp.GetInstance(),
	}
}

func Register() error {

	cli, err := api.NewClient(conf)
	if err != nil {
		return err
	}

	appName := gonfig.GetViper().GetString("name")
	env := gonfig.GetViper().GetString("env")
	addr := gonfig.GetViper().GetString("SERVER_ID")
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

	appName := gonfig.GetViper().GetString("name")
	if err := cli.Agent().ServiceDeregister(appName); err != nil {
		return err
	}

	log.Infof("consul deregister service: %s", appName)
	return nil
}
