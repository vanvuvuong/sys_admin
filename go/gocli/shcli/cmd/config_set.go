/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/gocli/shcli/library"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set permanant config file",
	Run: func(cmd *cobra.Command, args []string) {
		library.Pf("Not yet implement.\n")
		// if len(args) > 0 {
		// 	library.Log("Not support multiple config files. Please provide config file with flag `--config`")
		// }
		// configFile := rootCmd.PersistentFlags().Lookup("config")
		// if _, err := os.Stat(configFile.Value.String()); err != nil {
		// 	library.Log(err) // check file exists
		// }
		// configContent := library.BufferReadFile(configFile.Value.String())
		// configYaml := configStruct{}
		// err := yaml.Unmarshal([]byte(configContent), configYaml)
		// if err != nil {
		// 	library.Log(err)
		// }
		// viper.Setde
	},
}

func init() {
	configCmd.AddCommand(setCmd)
}
