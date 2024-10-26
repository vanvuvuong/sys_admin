/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/gocli/shcli/library"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get current config file path",
	Run: func(cmd *cobra.Command, args []string) {
		library.Pf("Current used config file: %s\n", viper.ConfigFileUsed())
	},
}

func init() {
	configCmd.AddCommand(getCmd)
}
