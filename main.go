/*
Copyright © 2022 Barnando Akbarto Hidayatullah <barnando13@gmail.com>

*/
package main

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/xxidbr9/rage/cmd"
)

func main() {
	type PackageJson struct {
		Version string `json:"version"`
	}
	packageJson := PackageJson{}

	viperConfig := viper.New()
	viperConfig.AddConfigPath(".")
	viperConfig.SetConfigType("json")
	viperConfig.SetConfigName("package")
	err := viperConfig.ReadInConfig()
	if err != nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	viper.AutomaticEnv()
	viperConfig.Unmarshal(&packageJson)

	version := packageJson.Version
	if version == "" {
		version = "0.1.0" // todo
	}
	cmd.Execute(version)
}
