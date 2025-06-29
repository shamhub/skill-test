package handlers

import (
	"os"
	"path/filepath"
	"strings"
)

func isPathValid(filePath string) bool {
	fileName := filepath.Base(filePath)
	if fileName == "." || fileName == "/" {
		return false
	}

	// check file extension pdf
	extension := filepath.Ext(fileName)
	isPDF := strings.ToLower(extension) == ".pdf"
	if !isPDF {
		return false
	}

	// Check if the folder exists
	folderPath := filepath.Dir(filePath)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		return false
	}

	return true
}
