package models

import (
	"io/ioutil"
	"path/filepath"
)

// ReadFile - read file and return string contents
func ReadFile(relativePath string) ([]byte, error) {
	absPath, absPathErr := filepath.Abs(relativePath)
	if absPathErr != nil {
		return nil, absPathErr
	}

	fileContents, readFileErr := ioutil.ReadFile(absPath)
	if readFileErr != nil {
		return nil, readFileErr
	}

	return fileContents, nil
}
