/*
Copyright © 2022 Barnando Akbarto Hidayatullah <barnando13@gmail.com>

*/
package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/xxidbr9/rage/shared/helpers"
)

var CompName string
var CompClassify string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new Components (atom|molecule|organism|template)",
	// TODO
	Long: ``,
	Run:  createCmdRunner,
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createCmd.Flags().StringVarP(&CompClassify, "component", "c", CompClassify, "Classify the component (atom|molecule|organism|template)")
	createCmd.Flags().StringVarP(&CompName, "name", "n", CompName, "Naming components")
}

func createCmdRunner(cmd *cobra.Command, args []string) {
	var componentName string
	var componentClassify ComponentsType

	if CompName != "" {
		componentName = CompName
	} else {
		if len(args) == 0 {
			log.Panicln(errors.New("require name"))
		} else {
			componentName = args[len(args)-1]
		}
	}

	if CompClassify != "" {
		componentClassify = ComponentsType(CompClassify)
	}

	if err := componentClassify.Validate(); err != nil {
		panic(fmt.Errorf("components classifier are required %s", err))
	}

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	storedLocation := filepath.Join(componentsOutDir, componentClassify.ToPlural(), componentName)

	// Create new Folder
	if err = helpers.CreateFolder(storedLocation); err != nil {
		panic(err)
	}

	if err = helpers.CreatePatternComponent(dir, storedLocation, componentName); err != nil {
		panic(err)
	}

}
