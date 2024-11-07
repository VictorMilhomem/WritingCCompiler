package main

import (
	"fmt"
	"unicode"
)

type TokenKind int

const (
	EOF TokenKind = iota
	ERROR
	ID
	NUMBERINT
	NUMBERFLOAT
	RPAREN
	LPAREN
	RBRACE
	LBRACE
	SEMICOLON
	MINUS
	PLUS
	DECREMENT
	INCREMENT
	BITWISE

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

var kindString = map[TokenKind]string{
	EOF:         "EOF",
	ERROR:       "ERROR",
	ID:          "ID",
	NUMBERINT:   "INTEGER",
	NUMBERFLOAT: "FLOAT",
	RPAREN:      "RPAREN",
	LPAREN:      "LPAREN",
	RBRACE:      "RBRACE",
	LBRACE:      "LBRACE",
	SEMICOLON:   "SEMICOLON",
	INT:         "KEYWORD",
	VOID:        "KEYWORD",
	RETURN:      "KEYWORD",
}

type Location struct {
	Start, End int
	Text, File string
}

type Token struct {
	Kind     TokenKind
	Text     string
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

func NewToken(kind TokenKind, text string, fullText string) Token {
	return Token{Kind: kind, Text: text, FullText: fullText}
}

func (t Token) String() string {
	return fmt.Sprintf("Token (kind: '%s', text: '%s')", kindString[t.Kind], t.Text)
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

func (l *Lexer) Tokenizer() []Token {
	for {
		l.Start = l.Pos

		if l.Pos >= len(l.Input) {
			l.Tokens = append(l.Tokens, Token{Kind: EOF})
			break
		}

		if !l.NextToken() {
			break
		}
	}

	return l.Tokens
}

func (l *Lexer) NextToken() bool {
	switch c := l.Peek(); c {
	case ' ', '\t', '\r':
	case '(':
		l.Tokens = append(l.Tokens, Token{Kind: LPAREN, Text: string(c)})
	case ')':
		l.Tokens = append(l.Tokens, Token{Kind: RPAREN, Text: string(c)})
	case '{':
		l.Tokens = append(l.Tokens, Token{Kind: LBRACE, Text: string(c)})
	case '}':
		l.Tokens = append(l.Tokens, Token{Kind: RBRACE, Text: string(c)})
	case ';':
		l.Tokens = append(l.Tokens, Token{Kind: SEMICOLON, Text: string(c)})
	case '~':
		l.Tokens = append(l.Tokens, Token{Kind: BITWISE, Text: string(c)})
	case '-':
		lookahead := l.LookAhead(1)
		switch lookahead {
		case '-':
			l.Tokens = append(l.Tokens, Token{Kind: DECREMENT, Text: string(c)})
			l.Advance(2)
		default:
			l.Tokens = append(l.Tokens, Token{Kind: MINUS, Text: string(c)})
		}
	case '+':
		lookahead := l.LookAhead(1)
		switch lookahead {
		case '+':
			l.Tokens = append(l.Tokens, Token{Kind: INCREMENT, Text: string(c)})
			l.Advance(2)
		default:
			l.Tokens = append(l.Tokens, Token{Kind: PLUS, Text: string(c)})
		}

	default:
		if unicode.IsLetter(c) {
			return l.Identifier()
		} else if unicode.IsDigit(c) {
			return l.Number()
		}
		//l.Tokens = append(l.Tokens, l.MakeToken(ERROR))
	}
	l.Pos++
	return true
}

func (l *Lexer) MakeToken(kind TokenKind) Token {
	text := l.Input[l.Start:l.Pos]

	token := NewToken(kind, text, text)
	return token
}

func (l *Lexer) Identifier() bool {
	l.Advance(1) // skip the first letter

	for !l.Eof() && IsIdSegment(l.Peek()) {
		l.Advance(1)
	}

	id := l.Input[l.Start:l.Pos]

	if keyword, ok := keywords[id]; ok {
		l.Tokens = append(l.Tokens, l.MakeToken(keyword))
	} else {

		l.Tokens = append(l.Tokens, l.MakeToken(ID))
	}
	return true
}

func (l *Lexer) Number() bool {
	l.Advance(1)
	for !l.Eof() && unicode.IsDigit(l.Peek()) {
		l.Advance(1)

		if !l.Eof() && l.Peek() == '.' {
			l.Advance(1)
			for !l.Eof() && unicode.IsDigit(l.Peek()) {
				l.Advance(1)
			}
			l.Tokens = append(l.Tokens, l.MakeToken(NUMBERFLOAT))
		}
	}
	l.Tokens = append(l.Tokens, l.MakeToken(NUMBERINT))
	return true
}

func (l *Lexer) Peek() rune {
	return rune(l.Input[l.Pos])
}

func (l *Lexer) Advance(n int) {
	l.Pos += n
}

func (l *Lexer) Eof() bool {
	return l.Pos >= len(l.Input)
}

func (l *Lexer) LookAhead(n int) rune {
	return rune(l.Input[l.Pos+n])
}

func IsIdSegment(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '.' || r == '\''
}
