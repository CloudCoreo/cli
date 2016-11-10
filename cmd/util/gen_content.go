package util

import (
	"os"
	"path/filepath"
)

// CreateFolder create folder
func CreateFolder(name, path string) error {
	if _, err := os.Stat(filepath.Join(path, name)); err != nil {
		// set default permissions
		os.Mkdir(filepath.Join(path, name), 0755)
	}

	return nil
}

// CreateFile create a file
func CreateFile(name, path, content string, override bool) error {
	if _, err := os.Stat(filepath.Join(path, name)); err != nil || override {
		f, _ := os.Create(filepath.Join(path, name))
		defer f.Close()

		if _, err := f.WriteString(content); err != nil {
			return err
		}
	}

	return nil
}
