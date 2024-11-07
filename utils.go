package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func GetDirectoryFiles(path string) []string {
	files, err := os.ReadDir(path)
	filesnames := []string{}
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		filesnames = append(filesnames, path+"/"+file.Name())
	}
	return filesnames
}

func GetFileSource(fpath string) string {
	bytes, err := ioutil.ReadFile(fpath)
	source := string(bytes)
	if err != nil {
		log.Fatalf("could not open file %s", fpath)
	}
	return source
}

func writeAssemblyToFile(assembly, outputFile string) error {
	dir := path.Dir(outputFile)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("could not create directory %s: %v", dir, err)
	}

	// Write the assembly string to the file
	if err := ioutil.WriteFile(outputFile, []byte(assembly), 0644); err != nil {
		return fmt.Errorf("could not write to file %s: %v", outputFile, err)
	}
	return nil
}
