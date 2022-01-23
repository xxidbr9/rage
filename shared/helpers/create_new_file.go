package helpers

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/aymerick/raymond"
	"github.com/xxidbr9/rage/shared/templates"
)

// Todo
func createJSXComponents(dir, storedFile, template, componentsName string) error {
	/*
		? TODO
		[ ] Add extension params
		[ ] Check if file are exist
	*/

	// Read Template
	// templateFile, err := ioutil.ReadFile(templateLocation)
	// if err != nil {
	// 	return err
	// }

	// Render
	ctx := map[string]string{
		"ComponentsName": componentsName,
	}

	result, err := raymond.Render(template, ctx)
	if err != nil {
		panic("Please report a bug :)")
	}

	// Write file
	perm := os.FileMode(0644)
	return ioutil.WriteFile(storedFile, []byte(result), perm)

}

// TODO
func CreateReadme() error {
	return nil
}

// Todo
func createIndexComponents(dir, storedLocation, template, componentName string) error {
	// templateLocation := filepath.Join(dir, "shared", "templates", "index_components.hbs")
	storedFile := filepath.Join(storedLocation, "index.tsx")

	return createJSXComponents(dir, storedFile, template, componentName)

}

// TODO
func CreatePatternComponent(dir, storedLocation, componentName string) error {
	// templateLocation := filepath.Join(dir, "shared", "templates", "components.hbs")
	storedFile := filepath.Join(storedLocation, componentName+".tsx")

	componentTemplate := templates.ComponentsJSX
	indexTemplate := templates.IndexJSX

	if err := createJSXComponents(dir, storedFile, componentTemplate, componentName); err != nil {
		return err
	}

	if err := createIndexComponents(dir, storedLocation, indexTemplate, componentName); err != nil {
		return err
	}

	return nil
}
