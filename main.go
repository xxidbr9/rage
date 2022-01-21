/*
Copyright © 2022 Barnando Akbarto Hidayatullah <barnando13@gmail.com>

*/
package main

import (
	"fmt"
	"io/ioutil"

	jsoniter "github.com/json-iterator/go"
	"github.com/xxidbr9/rage/cmd"
)

func main() {
	type PackageJson struct {
		Version string `json:"version"`
	}

	jsonFile, err := ioutil.ReadFile("./package.json")
	if err != nil {
		fmt.Println(err)
	}

	packageJson := PackageJson{}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(jsonFile, &packageJson)
	if err != nil {
		fmt.Println(err)
	}

	cmd.Execute(packageJson.Version)
}
