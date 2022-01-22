/*
Copyright © 2022 Barnando Akbarto Hidayatullah <barnando13@gmail.com>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type ComponentsType string

const (
	AtomComp     ComponentsType = "atom"
	MoleculeComp ComponentsType = "molecule"
	Organism     ComponentsType = "organism"
	Template     ComponentsType = "template"
)

func (c ComponentsType) Validate() error {
	switch c {
	case AtomComp:
	case MoleculeComp:
	case Organism:
	case Template:
		return nil
	default:
		return fmt.Errorf("select between (%s|%s|%s|%s)", AtomComp, MoleculeComp, Organism, Template)
	}
	return nil
}

func (c *ComponentsType) ToPlural() string {
	return string(*c) + "s"
}

func ToPlural(comp ComponentsType) string {
	return string(comp) + "s"
}

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
