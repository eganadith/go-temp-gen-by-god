package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	// Get a list of all files in the current directory
	files, err := ioutil.ReadDir(currentDir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// Loop through the files and convert .go files to .tmp files
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".go") {
			goFilePath := filepath.Join(currentDir, file.Name())
			tmpFilePath := filepath.Join(currentDir, strings.TrimSuffix(file.Name(), ".go")+".tmp")

			// Read the content of the .go file
			content, err := ioutil.ReadFile(goFilePath)
			if err != nil {
				fmt.Printf("Error reading %s: %v\n", goFilePath, err)
				continue
			}

			// Write the content to the .tmp file
			err = ioutil.WriteFile(tmpFilePath, content, 0644)
			if err != nil {
				fmt.Printf("Error writing to %s: %v\n", tmpFilePath, err)
				continue
			}

			fmt.Printf("Converted %s to %s\n", goFilePath, tmpFilePath)
		}
	}
}

git init
git add .
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/eganadith/go-temp-gen-by-god.git
git push -u origin main