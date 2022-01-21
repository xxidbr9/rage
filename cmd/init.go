/*
Copyright © 2022 Barnando Akbarto Hidayatullah <barnando13@gmail.com>

*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var componentsOutDir string = "./src/components"
var subFolder []string = []string{"atoms", "molecule", "organisms", "templates"}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initial Components Projects",
	// TODO
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		out := "Initialzing atomic components on"
		if componentsOutDir != "" {
			fmt.Printf("%s \"%s\"\n", out, componentsOutDir)
			err := createNewFolder(componentsOutDir)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("Successfully creating \"%s\"\n", componentsOutDir)
			for _, item := range subFolder {
				path := filepath.Join(componentsOutDir, item)
				createSubFolder(path)
			}
		} else {
			fmt.Printf("%s \"src/components\"", out)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&componentsOutDir, "out", "o", componentsOutDir, "location for components")
}

func createNewFolder(dir string) error {
	pwd, err := os.Getwd()
	location := filepath.Join(pwd, dir)
	err = os.MkdirAll(location, 0755)

	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		if _, err := os.Stat(location); os.IsNotExist(err) {
			os.Mkdir(location, 0755)
		}
	}
	if os.IsExist(err) {
		// check that the existing path is a directory
		info, err := os.Stat(dir)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return errors.New("already exists but is not a directory")
		}
		return nil
	}

	return err
}

func createSubFolder(dir string) error {
	pwd, err := os.Getwd()
	location := filepath.Join(pwd, dir)
	err = os.MkdirAll(location, 0755)

	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		if _, err := os.Stat(location); os.IsNotExist(err) {
			os.Mkdir(location, 0755)
		}
	}
	if os.IsExist(err) {
		// check that the existing path is a directory
		info, err := os.Stat(dir)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return errors.New("already exists but is not a directory")
		}
		return nil
	}

	return err
}
