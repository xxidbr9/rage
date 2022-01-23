/*
Copyright © 2022 Barnando Akbarto Hidayatullah <barnando13@gmail.com>

*/
package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/viper"
	"github.com/xxidbr9/rage/cmd"
)

func main() {
	type PackageJson struct {
		Version string `json:"version"`
	}
	var packageJson PackageJson

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

	var version string
	if packageJson.Version == "" {
		cmd := exec.Command("npm", "list", "-g", "@xxidbr9/rage", "version")
		versionRaw, err := cmd.CombinedOutput()
		if err != nil {
			panic(err)
		}
		versionArr := strings.Split(string(versionRaw), "@")
		version = strings.TrimSpace(versionArr[len(versionArr)-1])
	} else {
		version = packageJson.Version
	}

	cmd.Execute(version)
}
