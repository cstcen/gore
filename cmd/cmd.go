package cmd

import (
	"git.tenvine.cn/backend/gore/gonfig"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Example: "-n gdis -e sdev0"}

func init() {
	rootCmd.PersistentFlags().StringP("name", "n", "", "Application name")
	rootCmd.PersistentFlags().StringP("env", "e", "", "Environment name")
	rootCmd.PersistentFlags().StringP("port", "p", "8000", "Tcp port server listening on")
	rootCmd.PersistentFlags().StringP("consul", "c", "i-consul-${profile}.xk5.com", "Consul host or host:port")
	rootCmd.PersistentFlags().BoolP("log", "l", false, "Enable log")
	rootCmd.PersistentFlags().BoolP("host", "h", false, "暂时无用处，只是兼容这个参数")
	_ = gonfig.Instance().BindPFlags(rootCmd.PersistentFlags())
}

func GetInstance() *cobra.Command {
	return rootCmd
}
