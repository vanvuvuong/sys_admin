/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:     "user <UID 1>...<UID n>",
	Aliases: []string{"u"},
	Short:   "Export user(s) data",
	Long: `Export users data. Default option will export all users data to a CSV file.

Otherwise, please provide specific users ID to export data, or use flags like:
	-f users.csv: export user data based on users id CSV file (CSV file structure: user_id,username)
	-pf users.csv: export user data based on username regex pattern (CSV file structure: username_pattern)
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("user called")
		run(args)
	},
}

func init() {
	exportCmd.AddCommand(userCmd)

	// Here you will define your flags and configuration settings.
	userCmd.Flags().StringP("file", "f", "", "Get user(s) data based on CSV file")
}

func run(args []string) {
	fmt.Println(len(args))
	fmt.Println(viper.GetString("url"))
	fmt.Println(viper.GetString("path.user"))
	for key, value := range viper.GetStringMap("headers") {
		fmt.Println(key, value)
	}
}
