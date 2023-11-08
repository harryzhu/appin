/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var (
	osType         string
	PWD            string
	ADBPATH        string
	AndroidVersion int
	Package        string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "appin",
	Short: "appin [pin | unpin] [options]",
	Long:  ``,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	osType = strings.ToLower(runtime.GOOS)

	PWD, _ = os.Getwd()

	if osType == "windows" {
		ADBPATH = strings.Join([]string{PWD, "_3rd/bin", osType, "adb.exe"}, "/")
	} else {
		ADBPATH = strings.Join([]string{PWD, "_3rd/bin", osType, "adb"}, "/")
	}

	_, err := os.Stat(ADBPATH)
	if err != nil {
		log.Println("cannot find adb")
		ADBPATH = "adb"
	}

	fmt.Println(ADBPATH)

	rootCmd.PersistentFlags().StringVar(&Package, "pkg", "", "")
	rootCmd.PersistentFlags().IntVar(&AndroidVersion, "android-version", 0, "")

}
