package cmd

import (
	"fmt"
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

	processedContent, err := functionNamingPass(string(content))
	if err != nil {
		return fmt.Errorf("error in functionNamingPass: %v", err)
	}

	processedContent, err = semicolonRemovalPass(string(processedContent))
	if err != nil {
		return fmt.Errorf("error in semicolonRemovalPass: %v", err)
	}

	processedContent, err = returnTypePass(string(processedContent))
	if err != nil {
		return fmt.Errorf("error in returnTypePass: %v", err)
	}
	print(string(processedContent))
	processedContent, err = automaticTypePass(string(processedContent))
	if err != nil {
		return fmt.Errorf("error in automaticTypePass: %v", err)
	}

	goFilePath := filepath.Join(buildFolder, strings.TrimSuffix(filepath.Base(file), ".gone")+".go")
	return os.WriteFile(goFilePath, []byte(processedContent), 0644)
}

func functionNamingPass(content string) (string, error) {
	var processedContent string

	index := 0
	for index < len(content) {
		char, newIndex := getNextChar(content, index)
		index = newIndex

		if len(content)-index >= 8 && content[index:index+8] == "function" {
			processedContent += "func"
			index += 8
		} else {
			processedContent += string(char)
		}
	}

	return processedContent, nil
}

func semicolonRemovalPass(content string) (string, error) {
	var processedContent string

	index := 0
	for index < len(content) {
		char, newIndex := getNextChar(content, index)
		index = newIndex
		if char != ';' {
			processedContent += string(char)
		}
	}

	return processedContent, nil
}

func returnTypePass(content string) (string, error) {
	var processedContent string

	index := 0
	openBrace := false

	for index < len(content) {
		char, newIndex := getNextChar(content, index)
		index = newIndex

		if len(content)-index >= 2 && content[index:index+2] == "->" {
			openBraceIndex := index + 2
			for openBraceIndex < len(content) && content[openBraceIndex] != '{' {
				openBraceIndex++
			}
			commaIndex := index + 2
			for commaIndex < openBraceIndex && content[commaIndex] != ',' {
				commaIndex++
			}

			if commaIndex < openBraceIndex {
				processedContent += " ("
				openBrace = true
			} else {
				processedContent += ""
			}

			index += 2
		} else {
			if openBrace && content[index-1] == '{' {
				processedContent += ") "
				openBrace = false
			}
			processedContent += string(char)
		}
	}

	return processedContent, nil
}

func automaticTypePass(content string) (string, error) {
	var processedContent string
	index := 0
	letFlag := false
	for index < len(content) {
		char, newIndex := getNextChar(content, index)
		index = newIndex
		if len(content)-index >= 3 && content[index:index+3] == "let" {
			index += 3
			letFlag = true
		} else if len(content)-index >= 1 && content[index:index+1] == "=" && letFlag {
			processedContent += string(char)
			processedContent += ":"
		} else {
			processedContent += string(char)
		}
	}

	return processedContent, nil
}

func keywordFinder(content string, index int, keyword string) int {
	for index < len(content) && !strings.HasPrefix(content[index:], keyword) {
		index++
	}
	return index + 1
}

func getNextChar(content string, index int) (byte, int) {
	//fmt.Printf("getNextChar: Processing character '%c' at index %d\n", content[index], index)

	switch char := content[index]; char {
	case '\'':
		return content[index], keywordFinder(content, index, "'")
	case '"':
		return content[index], keywordFinder(content, index, "\"")
	case '/':
		if index < len(content)-1 {
			if content[index+1] == '/' {
				return '\n', keywordFinder(content, index, "\n")
			} else if content[index+1] == '*' {
				return content[index], keywordFinder(content, index, "*/") + 2
			}
		}
	}

	return content[index], index + 1
}
