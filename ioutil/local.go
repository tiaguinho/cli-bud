package ioutil

import (
	"log"
	"os"
	"path/filepath"
)

// LocalFile represents a file with the content
type LocalFile struct {
	Path     string
	Filename string
	Content  string
}

// ReadPath reads the path and returns the files
func ReadPath(path string) []LocalFile {
	finfo, err := os.Lstat(path)
	if err != nil {
		log.Fatal(err)
	}

	readed := make([]LocalFile, 0)
	if finfo.IsDir() {
		log.Println("Reading directory:", finfo.Name())
		readed = readDir(path)
	} else {
		log.Println("Reading file:", finfo.Name())
		readed = append(readed, readFile(path))
	}

	return readed
}

func readDir(path string) []LocalFile {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	readed := make([]LocalFile, 0)
	for _, file := range files {
		log.Println("Reading file:", file.Name())
		if file.IsDir() {
			log.Println(file.Name(), "is a directory. Skipping...")
			continue
		}

		readed = append(readed, readFile(filepath.Join(path, file.Name())))
	}

	return readed
}

func readFile(path string) LocalFile {
	// Read the file
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// Return the file metadata and content
	return LocalFile{
		Path:     path,
		Filename: filepath.Base(path),
		Content:  string(content),
	}
}
