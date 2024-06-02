package ioutil

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// AIFile represents a file with the response from ChatGPT
type AIFile struct {
	Original *LocalFile
	Suffix   string
	Content  string
}

// WriteAIFile writes the file to the output path
func WriteAIFile(file AIFile, output string) {
	// Get the new path to write the file
	var path string
	if file.Original != nil {
		path = file.Original.Path
		if output != "" {
			path = output + string(os.PathSeparator) + file.Original.Filename
		}
	} else {
		path = output
	}

	if file.Suffix != "" {
		ext := filepath.Ext(path)
		path = strings.Replace(path, ext, file.Suffix+ext, 1)
	}

	// Write the file
	log.Println("Writing:", path)
	err := os.WriteFile(path, []byte(file.Content), 0o644)
	if err != nil {
		log.Fatal(err)
	}
}
