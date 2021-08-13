package cmd

import (
	goreGin "git.tenvine.cn/backend/gore/gin"
	"git.tenvine.cn/backend/gore/gonfig"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Example: "-n gdis -e sdev0"}

func init() {
	rootCmd.PersistentFlags().StringP("name", "n", "", "Application name")
	rootCmd.PersistentFlags().StringP("env", "e", "", "Environment name")
	rootCmd.PersistentFlags().StringP("port", "p", "8000", "Tcp port server listening on")
	rootCmd.PersistentFlags().StringP("consul", "c", "https://i-consul-${profile}.xk5.com", "Tcp port server listening on")
	_ = gonfig.Instance().BindPFlags(rootCmd.PersistentFlags())
}

func New(setup func() error) *cobra.Command {
	rootCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		return setup()
	}
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		return goreGin.Startup()
	}
	return rootCmd
}
