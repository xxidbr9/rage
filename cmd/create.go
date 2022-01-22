/*
Copyright © 2022 Barnando Akbarto Hidayatullah <barnando13@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new Components (atom|molecule|organism|template)",
	// TODO
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},

}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
