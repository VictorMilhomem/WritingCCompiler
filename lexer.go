package main

import "fmt"

type TokenKind int

const (
	EOF TokenKind = iota
	ID
	CONSTANT
	RPAREN
	LPAREN
	RBRACE
	LBRACE
	SEMICOLON

	// keywords
	INT
	VOID
	RETURN
)

var keywords = map[string]TokenKind{
	"int":    INT,
	"void":   VOID,
	"return": RETURN,
}

type Location struct {
	Start, End int
	Text, File string
}

type Token struct {
	Kind     TokenKind
	Text     string
	Loc      Location
	FullText string
}

type Lexer struct {
	Filename string
	Input    string
	Tokens   []Token
	Pos      int
	Start    int
}

func NewLexer(filename, input string) *Lexer {
	return &Lexer{Filename: filename, Input: input}
}

func NewToken(kind TokenKind, text string, loc Location, fullText string) Token {
	return Token{Kind: kind, Text: text, Loc: loc, FullText: fullText}
}

func (t Token) String() string {
	return fmt.Sprintf("Token (kind: '%s', text: '%s')", t.Kind, t.Text)
}

func (l Location) LocStart() int {
	return l.Start
}

func (l Location) LocEnd() int {
	return l.End
}

func (l Location) LocText() string {
	return l.Text
}

func (l Location) LocFile() string {
	return l.File
}
