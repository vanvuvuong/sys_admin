/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/smirzaei/parallel"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		file := cmd.Flag("file")
		force := cmd.Flag("force")
		handleUpdate(args, file, force)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringP("file", "f", "", "update workflow based on CSV file.")
	updateCmd.Flags().Bool("force", false, "force to override old file data.")
}

func handleUpdate(args []string, flag ...*pflag.Flag) {
	index := 0
	parallel.ForEachLimit(args, maxConcurrency, func(param string) {
		index += 1
		log.Infof("Export connection %d/%d: %s.\n", index, len(args), param)
		handOneUpdate(param, flag...)
	})
}

func handOneUpdate(param string, flag ...*pflag.Flag) {
	fmt.Printf("update %s.\n", param)
}
