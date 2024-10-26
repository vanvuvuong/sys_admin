/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// type configStruct struct {
// 	url            string
// 	skip_ssl_check bool
// 	headers        struct {
// 		token    string
// 		tenantid int
// 	}
// 	path struct {
// 		user         string
// 		ifconnection string
// 	}
// }

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:     "config",
	Aliases: []string{"c"},
	Short:   "Check, create & set config file",
}

func init() {
	rootCmd.AddCommand(configCmd)
}
