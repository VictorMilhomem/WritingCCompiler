package main

import (
	"fmt"
	"io/ioutil"
	"log"
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

func main() {
	filepath := "C:\\Users\\victo\\GolandProjects\\RioLang\\tests\\simple.c"
	source := GetFileSource(filepath)
	lexer := NewLexer(filepath, source)
	tokens := lexer.Tokenizer()

	fmt.Printf("%v", tokens)
}
