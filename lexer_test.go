package main

import (
	"log"
	"reflect"
	"testing"
)

func normalizeTokens(tokens []Token) []Token {
	normalized := make([]Token, len(tokens))
	for i, token := range tokens {
		// Create a new Token with only kind and text fields
		normalized[i] = Token{
			Kind: token.Kind,
			Text: token.Text,
		}
	}
	return normalized
}

func TestLexer_Tokenizer_Valid(t *testing.T) {
	//filenames := GetFileSource("tests/chapter_1/valid")
	tests := []struct {
		name     string
		filepath string
		expected []Token
	}{
		{
			name:     "multi_digit",
			filepath: "tests/chapter_1/valid/multi_digit.c",
			expected: []Token{
				{Kind: INT, Text: "int"},
				{Kind: ID, Text: "main"},
				{Kind: LPAREN, Text: "("},
				{Kind: VOID, Text: "void"},
				{Kind: RPAREN, Text: ")"},
				{Kind: LBRACE, Text: "{"},
				{Kind: RETURN, Text: "return"},
				{Kind: NUMBERINT, Text: "100"},
				{Kind: SEMICOLON, Text: ";"},
				{Kind: RBRACE, Text: "}"},
				{Kind: EOF, Text: ""},
			},
		},
		{
			name:     "newlines",
			filepath: "tests/chapter_1/valid/newlines.c",
			expected: []Token{
				{Kind: INT, Text: "int"},
				{Kind: ID, Text: "main"},
				{Kind: LPAREN, Text: "("},
				{Kind: VOID, Text: "void"},
				{Kind: RPAREN, Text: ")"},
				{Kind: LBRACE, Text: "{"},
				{Kind: RETURN, Text: "return"},
				{Kind: NUMBERINT, Text: "0"},
				{Kind: SEMICOLON, Text: ";"},
				{Kind: RBRACE, Text: "}"},
				{Kind: EOF, Text: ""},
			},
		},
		{
			name:     "no_newlines",
			filepath: "tests/chapter_1/valid/no_newlines.c",
			expected: []Token{
				{Kind: INT, Text: "int"},
				{Kind: ID, Text: "main"},
				{Kind: LPAREN, Text: "("},
				{Kind: VOID, Text: "void"},
				{Kind: RPAREN, Text: ")"},
				{Kind: LBRACE, Text: "{"},
				{Kind: RETURN, Text: "return"},
				{Kind: NUMBERINT, Text: "0"},
				{Kind: SEMICOLON, Text: ";"},
				{Kind: RBRACE, Text: "}"},
				{Kind: EOF, Text: ""},
			},
		},
		{
			name:     "return_0",
			filepath: "tests/chapter_1/valid/return_0.c",
			expected: []Token{
				{Kind: INT, Text: "int"},
				{Kind: ID, Text: "main"},
				{Kind: LPAREN, Text: "("},
				{Kind: VOID, Text: "void"},
				{Kind: RPAREN, Text: ")"},
				{Kind: LBRACE, Text: "{"},
				{Kind: RETURN, Text: "return"},
				{Kind: NUMBERINT, Text: "0"},
				{Kind: SEMICOLON, Text: ";"},
				{Kind: RBRACE, Text: "}"},
				{Kind: EOF, Text: ""},
			},
		},
		{
			name:     "return_2",
			filepath: "tests/chapter_1/valid/return_2.c",
			expected: []Token{
				{Kind: INT, Text: "int"},
				{Kind: ID, Text: "main"},
				{Kind: LPAREN, Text: "("},
				{Kind: VOID, Text: "void"},
				{Kind: RPAREN, Text: ")"},
				{Kind: LBRACE, Text: "{"},
				{Kind: RETURN, Text: "return"},
				{Kind: NUMBERINT, Text: "2"},
				{Kind: SEMICOLON, Text: ";"},
				{Kind: RBRACE, Text: "}"},
				{Kind: EOF, Text: ""},
			},
		},
		{
			name:     "spaces",
			filepath: "tests/chapter_1/valid/spaces.c",
			expected: []Token{
				{Kind: INT, Text: "int"},
				{Kind: ID, Text: "main"},
				{Kind: LPAREN, Text: "("},
				{Kind: VOID, Text: "void"},
				{Kind: RPAREN, Text: ")"},
				{Kind: LBRACE, Text: "{"},
				{Kind: RETURN, Text: "return"},
				{Kind: NUMBERINT, Text: "0"},
				{Kind: SEMICOLON, Text: ";"},
				{Kind: RBRACE, Text: "}"},
				{Kind: EOF, Text: ""},
			},
		},
		{
			name:     "tabs",
			filepath: "tests/chapter_1/valid/tabs.c",
			expected: []Token{
				{Kind: INT, Text: "int"},
				{Kind: ID, Text: "main"},
				{Kind: LPAREN, Text: "("},
				{Kind: VOID, Text: "void"},
				{Kind: RPAREN, Text: ")"},
				{Kind: LBRACE, Text: "{"},
				{Kind: RETURN, Text: "return"},
				{Kind: NUMBERINT, Text: "0"},
				{Kind: SEMICOLON, Text: ";"},
				{Kind: RBRACE, Text: "}"},
				{Kind: EOF, Text: ""},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Read the file content
			log.Println(tt.filepath)
			content := GetFileSource(tt.filepath)
			// Initialize the lexer with the file content as input
			lexer := &Lexer{
				Input: content,
			}

			// Tokenize the input
			tokens := lexer.Tokenizer()

			normalizedGot := normalizeTokens(tokens)
			normalizedExpected := normalizeTokens(tt.expected)

			// Compare the normalized tokens
			if !reflect.DeepEqual(normalizedGot, normalizedExpected) {
				t.Errorf("got %v, want %v", normalizedGot, normalizedExpected)
			}
		})
	}

}
