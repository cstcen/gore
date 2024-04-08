package gonfig

import (
	"bytes"
	"fmt"
	crypt "github.com/bketelsen/crypt/config"
	"github.com/cstcen/gore/common"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"os"
	"path/filepath"
	"strings"
)

var (
	vp = viper.New()

	storeKeyFns = []func(env, appName string) string{
		func(env, appName string) string {
			return fmt.Sprintf("config/go/%s/data", appName)
		},
		func(env, appName string) string {
			return fmt.Sprintf("config/go/%s,%s/data", appName, env)
		},
	}
)

func init() {
	vp.SetConfigType("yaml")
	vp.Set("gore.path", "config")
	vp.Set("gore.filename", "config.yml")
	vp.Set("gore.filenameEnv", "config-${profile}.yml")
	vp.AutomaticEnv()
}

func Instance() *viper.Viper {
	return vp
}

// Setup 读取顺序为：
// 假如env=dev, appName=gore
// 先读取consul K/V store
// 1. config/config.yml
// 2. config/config-dev.yml
// 再读取项目内配置文件
// 1. config/application/data
// 2. config/application,dev/data
// 3. config/gore/data
// 4. config/gore,dev/data
// 越靠后，配置优先级越高
func Setup() error {

	env := vp.GetString("env")
	appName := vp.GetString("name")
	replacer := strings.NewReplacer("${profile}", env, "${application}", appName)
	placeholder(replacer)

	if err := unmarshalConfigCustom(); err != nil {
		return err
	}
	if len(env) > 0 {
		if err := unmarshalConfigCustomEnv(env); err != nil {
			return err
		}
	}

	endpoint := vp.GetString("consul")
	name := "application"
	if err := readRemoteConfig(env, name, endpoint); err != nil {
		return err
	}
	if len(appName) > 0 {
		if err := readRemoteConfig(env, appName, endpoint); err != nil {
			return err
		}
	}

	placeholder(replacer)

	return nil
}

func SetDefaultCmdConfig(opt *common.Args) {
	vp.Set("name", opt.Name)
	vp.Set("env", opt.Env)
	vp.Set("consul", opt.Consul)
	vp.Set("log", opt.Log)
}

func placeholder(replacer *strings.Replacer) {
	for _, key := range vp.AllKeys() {
		val := vp.GetString(key)
		if len(val) == 0 {
			continue
		}
		val = replacer.Replace(val)
		vp.Set(key, val)
	}
}

func readRemoteConfig(env string, appName string, endpoint string) error {
	var cm crypt.ConfigManager
	var err error

	cm, err = crypt.NewStandardConsulConfigManager([]string{endpoint})
	if err != nil {
		return err
	}

	for _, fn := range storeKeyFns {
		path := fn(env, appName)
		b, err := cm.Get(path)
		if err != nil {
			continue
		}
		if err := vp.MergeConfig(bytes.NewReader(b)); err != nil {
			return err
		}
	}
	return nil
}

func unmarshalConfigCustom() error {
	return unmarshal(filepath.Join(vp.GetString("gore.path"), vp.GetString("gore.filename")))
}

func unmarshalConfigCustomEnv(env string) error {
	return unmarshal(filepath.Join(vp.GetString("gore.path"), vp.GetString("gore.filenameEnv")))
}

func unmarshal(filename string) error {
	yml, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	if err := vp.MergeConfig(bytes.NewBuffer(yml)); err != nil {
		return err
	}

	return nil
}
