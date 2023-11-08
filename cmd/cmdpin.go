/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	//"fmt"
	//"log"

	"github.com/spf13/cobra"
)

// pinCmd represents the pin command
var pinCmd = &cobra.Command{
	Use:   "pin",
	Short: " ",
	Long:  ` `,
	Run: func(cmd *cobra.Command, args []string) {

		if AndroidVersion == 0 {
			AndroidVersion = AdbDetectAndroidVersion()
		}

		if Package != "" {
			AdbExec(Package, "pin")
		}

	},
}

func init() {
	rootCmd.AddCommand(pinCmd)

}
