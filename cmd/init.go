/*
Copyright © 2022 Barnando Akbarto Hidayatullah <barnando13@gmail.com>

*/
package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/xxidbr9/rage/shared/helpers"
)

// initCmd represents the init command
var componentsOutDir string = "./src/components"
var subFolder []string = []string{ToPlural(AtomComp), ToPlural(MoleculeComp), ToPlural(Organism), ToPlural(Template)}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initial Components Projects",
	// TODO
	Long: ``,
	Run:  initCmdRunner,
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&componentsOutDir, "out", "o", componentsOutDir, "location for components")
}

func initCmdRunner(cmd *cobra.Command, args []string) {
	var h _handler
	out := "Initialzing atomic components on"
	if componentsOutDir != "" {
		fmt.Printf("%s \"%s\"\n", out, componentsOutDir)
		err := h.createNewFolder(componentsOutDir)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Successfully creating \"%s\"\n", componentsOutDir)
		for _, item := range subFolder {
			path := filepath.Join(componentsOutDir, item)
			h.createNewFolder(path)
		}
	} else {
		fmt.Printf("%s \"src/components\"", out)
	}
}

type _handler struct{}

func (h _handler) createNewFolder(dir string) error {
	return helpers.CreateFolder(dir)
}

// TODO
func (h _handler) generateInitFile() {}

// TODO
func (h _handler) checkConfig() {}
