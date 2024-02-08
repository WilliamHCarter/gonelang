package cmd

import (
	"os"
	"path/filepath"
	"strings"
)

const buildFolder = "gone-build"

// Generate scans for '.gone' files and converts them to '.go' files in the build folder
func Generate() error {
	if err := createBuildFolder(); err != nil {
		return err
	}

	files, err := findGoneFiles()
	if err != nil {
		return err
	}

	for _, file := range files {
		if err := processFile(file); err != nil {
			return err
		}
	}

	return nil
}

func createBuildFolder() error {
	if err := os.RemoveAll(buildFolder); err != nil {
		return err
	}

	return os.MkdirAll(buildFolder, 0755)
}

func findGoneFiles() ([]string, error) {
	var goneFiles []string

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(info.Name(), ".gone") {
			goneFiles = append(goneFiles, path)
		}
		return nil
	})

	return goneFiles, err
}

func processFile(file string) error {
	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	// Add 'println("testing")' to the end of the file
	content = append(content, []byte("\nprintln(\"testing\")\n")...)

	goFilePath := filepath.Join(buildFolder, strings.TrimSuffix(filepath.Base(file), ".gone")+".go")
	return os.WriteFile(goFilePath, content, 0644)
}
