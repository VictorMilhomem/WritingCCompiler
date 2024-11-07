package main

import (
	"log"
	"strconv"
)

type Parser struct {
	Tokens  []Token
	Current int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{tokens, 0}
}

func (p *Parser) Parse() AST {
	program := p.ParseProgram()
	p.Expect(EOF)
	return AST{program}
}

func (p *Parser) ParseProgram() Program {
	p.Expect(INT)
	function := p.ParseFunction()
	return Program{function}

}

func (p *Parser) ParseFunction() Function {
	p.Expect(ID)
	name := p.GetPreviousToken().Text
	p.Expect(LPAREN)
	//parse the function arguments
	p.Expect(VOID)

	p.Expect(RPAREN)

	p.Expect(LBRACE)
	body := p.ParseStatement()
	p.Expect(RBRACE)

	return Function{
		name,
		body,
	}
}

func (p *Parser) ParseStatement() Statement {
	p.Expect(RETURN)
	exp := p.ParseExpression()
	p.Expect(SEMICOLON)
	return Return{exp}

}

func (p *Parser) ParseExpression() Expression {
	p.Expect(NUMBERINT)
	token := p.GetPreviousToken()
	n, err := strconv.Atoi(token.Text)
	if err != nil {
		log.Fatal("Error Parsing Integer")
	}
	return NumberInteger{n}
}

func (p *Parser) Expect(expected TokenKind) {
	current := p.TakeToken().Kind
	if current != expected {

		log.Fatalf("Syntax Error %s", p.GetPreviousToken())
	}
}

func (p *Parser) TakeToken() Token {
	if p.Current < len(p.Tokens) {
		token := p.Tokens[p.Current]
		p.Current++
		return token
	}
	return Token{} // Return an empty token if out of bounds.
}

func (p *Parser) GetCurrentToken() Token {
	if p.Current < len(p.Tokens) {
		token := p.Tokens[p.Current]
		return token
	}
	return Token{} // Return an empty token if out of bounds.
}

func (p *Parser) GetPreviousToken() Token {
	if p.Current < len(p.Tokens) {
		token := p.Tokens[p.Current-1]
		return token
	}
	return Token{} // Return an empty token if out of bounds.
}
