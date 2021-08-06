package cmd

import (
	goreGin "git.tenvine.cn/backend/gore/gin"
	"git.tenvine.cn/backend/gore/gonfig"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Example: "-e sdev0 -n gdis -p 8080"}

func init() {
	rootCmd.PersistentFlags().StringP("name", "n", "", "Application name")
	rootCmd.PersistentFlags().StringP("env", "e", "", "Environment name")
	rootCmd.PersistentFlags().StringP("port", "p", "", "Tcp port server listening on")
	_ = gonfig.GetViper().BindPFlags(rootCmd.PersistentFlags())
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
