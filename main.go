package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func GetFileSource(fpath string) string {
	bytes, err := ioutil.ReadFile(path.Base(fpath))
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

func main() {
	filepath := "C:\\Users\\victo\\GolandProjects\\RioLang\\tests\\simple.c"
	source := GetFileSource(filepath)
	lexer := NewLexer(filepath, source)
	tokens := lexer.Tokenizer()
	parser := NewParser(tokens)
	parsedProgram := parser.Parse()
	assembly := generateAssembly(parsedProgram)
	writeAssemblyToFile(assembly, "C:\\Users\\victo\\GolandProjects\\RioLang\\tests\\simple_gen.s")
	// debug printing the ast
	PrintJson(parsedProgram)
	fmt.Printf(assembly)
}
