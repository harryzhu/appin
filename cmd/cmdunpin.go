/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	//"fmt"

	"github.com/spf13/cobra"
)

// unpinCmd represents the unpin command
var unpinCmd = &cobra.Command{
	Use:   "unpin",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if AndroidVersion == 0 {
			AndroidVersion = AdbDetectAndroidVersion()
		}

		if Package != "" {
			AdbExec(Package, "unpin")
		}

	},
}

func init() {
	rootCmd.AddCommand(unpinCmd)

}
