package consul

import (
	"github.com/cstcen/gore/gonfig"
	goreHttp "github.com/cstcen/gore/http"
	"github.com/cstcen/gore/log"
	"github.com/hashicorp/consul/api"
)

var (
	conf *api.Config
)

func SetConfig(cfg *api.Config) {
	conf = cfg
}

func SetupDefault() error {
	url := gonfig.Instance().GetString("gore.consul.discovery.host")
	conf = &api.Config{
		Address:    url,
		HttpClient: goreHttp.Instance(),
	}
	return nil
}

func Register() error {
	if !Enable() {
		return nil
	}

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
	if !Enable() {
		return nil
	}

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
