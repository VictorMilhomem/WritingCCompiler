package main

import (
	"fmt"
)

func main() {
	filepath := "tests\\simple.c"
	source := GetFileSource(filepath)
	lexer := NewLexer(filepath, source)
	tokens := lexer.Tokenizer()
	parser := NewParser(tokens)
	ast := parser.Parse()
	codeGen := CodeGenerator{}
	assembly := codeGen.Generate(ast)
	writeAssemblyToFile(assembly, "tests\\simple_gen.s")
	// debug printing the ast
	PrintJson(ast)
	fmt.Printf(assembly)
}
