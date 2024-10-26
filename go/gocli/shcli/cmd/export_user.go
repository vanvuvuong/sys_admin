/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"io"
	"strconv"

	"github.com/gocli/shcli/library"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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
	-p users.csv: export user data based on username regex pattern (CSV file structure: username_pattern)
`,
	Run: func(cmd *cobra.Command, args []string) {
		file := cmd.Flag("file")
		patternFile := cmd.Flag("pattern-file")
		run(args, file, patternFile)
	},
}

func init() {
	exportCmd.AddCommand(userCmd)

	// Here you will define your flags and configuration settings.
	userCmd.Flags().StringP("file", "f", "", "Get user(s) data based on user ID(s) CSV file")
	userCmd.Flags().StringP("pattern-file", "p", "", "Get user(s) data based on username pattern(s) CSV file")
}

func run(args []string, file, patternFile *pflag.Flag) {
	var (
		err       error
		url       = viper.GetString("url")
		user_path = viper.GetString("path.user")
		headers   = viper.GetStringMapString("headers")
	)
	based_url := library.Sf("%s/%s", url, user_path)
	// unhappy path
	if file.Changed && patternFile.Changed {
		library.Log("Error, please use one flag EITHER `-f/--file` OR `-p/--pattern-file`.")
	}
	if file.Changed && len(args) > 0 {
		library.Log("Error, please use EITHER user id list OR flag `-f/--file`.")
	}
	if patternFile.Changed && len(args) > 0 {
		library.Log("Error, please use EITHER user id list OR flag `-p/--pattern-file`")
	}

	// happy path
	if len(args) == 0 && !file.Changed && !patternFile.Changed {
		library.Pl("Get all user data...")
		err = getUserData(based_url, "all_users.json", headers) // default option will get all the userdata
	}
	if len(args) == 0 && file != nil {
		library.Pl("Get user data by CSV file...")
		err = getUserByCSVId(based_url, file.Value.String(), headers)
	} else if len(args) == 0 && patternFile != nil {
		// run get user by csv
	} else {
		for _, uid := range args {
			base_url_uid := library.Sf("%s/%s", based_url, uid)
			err := getUserData(base_url_uid, library.Sf("%s.json", uid), headers) // default option will get all the userdata
			if err != nil {
				library.Pf("When get user:", uid, "- Error:", err)
			}
		}
	}
	if err != nil {
		library.Log(err)
	}
	library.Pl("...Done get users data")
}

func getUserData(url, filename string, headers map[string]string) error {
	var err error
	resp, err := library.Request("GET", headers, []byte{}, url)
	if err != nil {
		return library.Err("Error to get user data: %w\n", err)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return library.Err("Error convert to string: %w", err)
	}
	err = library.BufferWriteFile(library.Sf("%s", filename), string(respBody))
	defer resp.Body.Close()
	return err
}

func getUserByCSVId(url, filename string, headers map[string]string) error {
	var err error
	csvData, err := library.DecodeCsv(filename)
	if err != nil {
		return library.Err("%w", err)
	} else if len(csvData) <= 1 {
		return library.Err("Error empty CSV file")
	}
	for index, row := range csvData {
		if len(row) != 2 {
			err = library.Err("Invalid CSV format file. Please check `-h/--help` command to see the CSV format.")
			break
		}
		if index == 0 {
			continue
		}
		uid, err := strconv.ParseUint(row[0], 0o0, 64)
		if err != nil {
			library.Pf("Invalid user ID. Please recheck the CSV file on row:%d.\n", index)
		}
		newUrl := library.Sf("%s/%d", url, uid)
		err = getUserData(newUrl, library.Sf("%d.json", uid), headers)
		library.Pl(err)
	}
	return err
}
