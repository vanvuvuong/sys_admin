/*
Copyright © 2024 Tran Dinh Dong

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"os"
	"strings"

	"github.com/gocli/shcli/library"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var (
	cfgFile string
	dryRun  bool
	rootCmd = &cobra.Command{
		Use:     "shcli",
		Short:   "CLI tool for fast get, update or create SH data",
		Version: "v0.0.1",
	}
)

// init function: will run first before any function in this file
func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (no default option)")
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "see how current command effects to the system")
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("SHCLI") // prefix SHCLI_XXX
	viper.SetConfigType("yml")
	viper.SetConfigFile("config_sample")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	versionTemplate := `{{printf "%s: %s - version: %s\n" .Name .Short .Version}}`
	rootCmd.SetVersionTemplate(versionTemplate)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// viper configuration file initilization
func initConfig() {
	if cfgFile == "" {
		library.Log("Not found config file.\n\nPlease run `shcli --config file.yml [subcommand]` or `shcli config --config file.yml` to set it.\nRun `shcli -h` for more detail\n")
		os.Exit(1)
	}
	if _, err := os.Stat(cfgFile); err != nil {
		library.Log(err) // check file exists
	}
	// Use config file from the flag.
	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv() // read in environment variables in config file
	viper.SetConfigFile(rootCmd.PersistentFlags().Lookup("config").Value.String())
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		library.Log("Using config file:", viper.ConfigFileUsed(), "- Error:", err)
	}
}
