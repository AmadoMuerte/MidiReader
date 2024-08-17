package getFiles

import (
	"os"
	"strings"
)

var fileNames []string

func GetFiles(path string, filter string) ([]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return fileNames, nil
	}

	var filteredFileNames []string
	for _, file := range files {
		if containsMidi(file.Name(), filter) {
			filteredFileNames = append(filteredFileNames, file.Name())
		}
	}

	return filteredFileNames, nil
}

func containsMidi(name string, filter string) bool {
	if strings.Contains(name, filter) {
		return true
	}

	return false
}
