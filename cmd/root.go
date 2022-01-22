/*
Copyright © 2022 Barnando Akbarto Hidayatullah <barnando13@gmail.com>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rage",
	Short: "React Atomic Generator",
	Long: `React Atomic Generator
My daily routine implement Atomic Pattern to the react app
i was doing that repeatedly and i want to make my life easier
so the reason to me create this package is for help me work every day`,
}

func Execute(version string) {
	rootCmd.Version = version
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Show Version")
}
