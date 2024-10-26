/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"io"

	"github.com/gocli/shcli/library"
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
		run(args)
	},
}

func init() {
	exportCmd.AddCommand(userCmd)

	// Here you will define your flags and configuration settings.
	userCmd.Flags().StringP("file", "f", "", "Get user(s) data based on user ID(s) CSV file")
	userCmd.Flags().StringP("pattern-file", "p", "", "Get user(s) data based on username pattern(s) CSV file")
}

func run(args []string) {
	var (
		url       = viper.GetString("url")
		user_path = viper.GetString("path.user")
		headers   = viper.GetStringMapString("headers")
	)
	based_url := library.Sf("%s/%s", url, user_path)
	library.Sf("Break here")
	file := exportCmd.Flags().Lookup("file")
	patternFile := exportCmd.Flags().Lookup("pattern-file")
	// unhappy path
	if file != nil && patternFile != nil {
		library.Log("Error, please use one flag EITHER `-f/--file` OR `-pf/--pattern-file`.")
	}
	if file != nil && len(args) > 0 {
		library.Log("Error, please use EITHER user id list OR flag `-f/--file`.")
	}
	if patternFile != nil && len(args) > 0 {
		library.Log("Error, please use EITHER user id list OR flag `-pf/--pattern-file`")
	}

	// happy path
	if len(args) == 0 && file == nil && patternFile == nil {
		library.Sf("Get all user data...")
		err := getUserData(based_url, "all_users.json", headers) // default option will get all the userdata
		if err != nil {
			library.Log(err)
		}
		library.Pf("...Done get all users data")
	}
	if len(args) == 0 && file != nil {
		// run get user by csv
	} else if len(args) == 0 && patternFile != nil {
		// run get user by csv
	} else {
		for _, uid := range args {
			base_url_uid := library.Sf("%s/%s", based_url, uid)
			err := getUserData(base_url_uid, library.Sf("%s.json", uid), headers) // default option will get all the userdata
			if err != nil {
				library.Sf("When get user:", uid, "- Error:", err)
			}
		}
	}
}

func getUserData(url, filename string, headers map[string]string) error {
	var err error
	resp, err := library.Request("GET", headers, []byte{}, url)
	defer resp.Body.Close()
	if err != nil {
		return library.Err("rror to get user data: %w", err)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return library.Err("Error convert to string: %w", err)
	}
	err = library.BufferWriteFile(library.Sf("%s", filename), string(respBody))
	return err
}

func getUserByCSVId() {
}
