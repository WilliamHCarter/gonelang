package cmd

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func Compile() {
	// Ensure the build directory exists.
	buildDir := "build"
	if _, err := os.Stat(buildDir); os.IsNotExist(err) {
		log.Fatalf("Error: The '%s' directory does not exist.", buildDir)
	}

	// Get the name of the Go file being compiled.
	goFileName := getGoFileName(buildDir)
	binaryName := removeExtension(goFileName)

	// Run 'go build' on the build directory.
	cmd := exec.Command("go", "build", "-o", binaryName)
	cmd.Dir = buildDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error building: %v\n%s", err, output)
	}

	// Move the binary back to the root folder.
	err = os.Rename(filepath.Join(buildDir, binaryName), binaryName)
	if err != nil {
		log.Fatalf("Error moving binary: %v", err)
	}

	// Remove the build directory.
	//if err := os.RemoveAll(buildDir); err != nil {
	//	log.Fatalf("Error removing build directory: %v", err)
	//}

	log.Printf("Compilation successful. Binary saved to: %s", binaryName)
}

func getGoFileName(buildDir string) string {
	goFiles, err := filepath.Glob(filepath.Join(buildDir, "*.go"))
	if err != nil {
		log.Fatalf("Error getting Go files: %v", err)
	}

	if len(goFiles) == 0 {
		log.Fatal("Error: No Go files found in the build directory.")
	}

	return filepath.Base(goFiles[0])
}

func removeExtension(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}
