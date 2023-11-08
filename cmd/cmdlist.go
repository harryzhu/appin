/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "",
	Long:  `list packages(3rd-party) and save into packages_3.txt`,
	Run: func(cmd *cobra.Command, args []string) {
		AdbList()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
