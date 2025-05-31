/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gocli/library"
	"github.com/op/go-logging"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var (
	verbose        bool
	file1          logging.Backend
	maxConcurrency = 1
	timestamp      = time.Now().Format(time.DateOnly)
	log            = logging.MustGetLogger("shcli")
	logFileName    = library.Sf("logs/shcli-%s.log", timestamp)
	format         = logging.MustStringFormatter(
		`%{color}%{time:15:04:0.00} ▶ %{level:.6s}%{color:reset} %{message}`,
	)
	fileFormat = logging.MustStringFormatter(
		`%{time:15:04:0.00} ▶ %{level:.6s} %{shortfile} %{callpath} %{id} %{message}`,
	)
	rootCmd = &cobra.Command{
		Use:   "gocli",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Version: "v0.0.1",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			sslCheck := viper.GetBool("ssl_check")
			if !sslCheck {
				log.Warning("Skip SSL checking.")
			}
			read := bufio.NewReader(os.Stdin)
			fmt.Fprintf(os.Stderr, "Confirm running? (y/n):\n>>> ")
			confirm, _ := read.ReadString('\n')
			confirm = strings.TrimSpace(confirm)
			confirm = strings.ToLower(confirm)
			if !strings.HasPrefix(confirm, "y") {
				os.Exit(1)
			}
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	// configure log
	logLevel := logging.INFO
	if verbose {
		logLevel = logging.DEBUG
	}
	std1 := logging.NewLogBackend(os.Stdout, "", 0)
	std1Leveled := logging.AddModuleLevel(std1)
	std1Leveled.SetLevel(logLevel, "")
	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err == nil {
		file1 = logging.NewLogBackend(file, "", 0)
	} else {
		log.Fatal("Failed to write log file: ", logFileName)
	}
	file1Formatter := logging.NewBackendFormatter(file1, fileFormat)
	logging.SetBackend(std1Leveled, file1Formatter)
	logging.SetFormatter(format)

	// handle multiple config
	replacer := strings.NewReplacer("-", "_")
	os.Mkdir("config", 0o755)
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("SHCLI")
	viper.BindEnv("env")
	viper.BindEnv("mode")
	viper.BindEnv("source_folder")
	viper.BindEnv("jfrog_env")
	if os.Getenv("SHCLI_ENV") != "" {
		configPath := library.Sf("config/config_%s.yml", os.Getenv("SHCLI_ENV"))
		if _, err := os.Stat(configPath); err != nil {
			genDefaultConfig(configPath)
		}
		viper.SetConfigFile(configPath)
	} else {
		if _, err := os.Stat("config/config.yml"); err != nil {
			genDefaultConfig("config/config.yml")
		}
		viper.SetConfigFile("config/config.yml")
	}
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error: '%v'. Please create '%s' config file.\n", err, viper.ConfigFileUsed())
	}
	if viper.GetString("env") != "" {
		fmt.Print("env: ", viper.GetString("env"))
	}
	if viper.GetInt("parrallel_job") != 0 {
		maxConcurrency = viper.GetInt("parrallel_job")
	}
}

func genDefaultConfig(filePath string) {
	yamlContent := `delay_in_request: 450
headers:
  accept: '*/*'
  authorization: Bearer token
  tenantid: 1
  user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0
parrallel_job: 3
`
	library.BufferWriteFile(filePath, yamlContent)
}
