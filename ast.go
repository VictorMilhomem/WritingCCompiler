package main

import (
	"encoding/json"
	"fmt"
)

type AST struct {
	Program `json:"program"`
}

type Program struct {
	Function `json:"function"`
}

type Function struct {
	Name string    `json:"name"`
	Body Statement `json:"body"`
}

type Statement interface {
	isStatement()
}

type Return struct {
	Expression `json:"expression"`
}

func (r Return) isStatement() {}

type Expression interface {
	isExpression()
}

type NumberInteger struct {
	Value int `json:"value"`
}

func (n NumberInteger) isExpression() {}

func PrintJson(ast AST) {
	jsonData, err := json.MarshalIndent(ast, "", "  ")
	if err != nil {
		fmt.Println("Error serializing to JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}
