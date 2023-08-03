package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func convertAndCopyGoFiles(srcDir, dstDir string) {
	// Read the list of files and directories inside the source directory
	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		fmt.Printf("Error reading directory %s: %v\n", srcDir, err)
		return
	}

	// Create the destination directory if it doesn't exist
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		fmt.Printf("Error creating directory %s: %v\n", dstDir, err)
		return
	}

	// Loop through the files and directories in the source directory
	for _, file := range files {
		srcFilePath := filepath.Join(srcDir, file.Name())
		dstFilePath := filepath.Join(dstDir, file.Name())

		if file.IsDir() {
			// Recursively process subdirectories
			convertAndCopyGoFiles(srcFilePath, dstFilePath)
		} else if strings.HasSuffix(file.Name(), ".go") {
			// Read the content of the .go file
			content, err := ioutil.ReadFile(srcFilePath)
			if err != nil {
				fmt.Printf("Error reading %s: %v\n", srcFilePath, err)
				continue
			}

			// Write the content to the .tmp file in the destination directory
			tmpFilePath := strings.TrimSuffix(dstFilePath, ".go") + ".tmp"
			if err := ioutil.WriteFile(tmpFilePath, content, 0644); err != nil {
				fmt.Printf("Error writing to %s: %v\n", tmpFilePath, err)
				continue
			}
			fmt.Printf("Converted %s to %s\n", srcFilePath, tmpFilePath)
		}
	}
}

func main() {
	srcBaseDir := "E:/Evolza/Go Temp Convert/Gofiles"
	dstBaseDir := "E:/Evolza/Go Temp Convert/NewTemp" // Change this to the destination path where you want to create the new folder structure with .tmp files

	convertAndCopyGoFiles(srcBaseDir, dstBaseDir)
}
