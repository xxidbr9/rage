package helpers

import (
	"errors"
	"os"
	"path/filepath"
)

func CreateFolder(dir string) error {
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
