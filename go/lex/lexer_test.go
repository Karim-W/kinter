package lex

import (
	"testing"

	"github.com/karim-w/kinter/go/tokens"
)

func TestToken(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		expectedType    tokens.Token
		expectedLiteral string
	}{
		{tokens.ASSGIN, "="},
		{tokens.EQUAL, "+"},
		{tokens.LPARENT, "("},
		{tokens.RPARENT, ")"},
		{tokens.LSQUIGL, "{"},
		{tokens.RSQUIGL, "}"},
		{tokens.COMMA, ","},
		{tokens.SEMICOLON, ";"},
		{tokens.EOF, ""},
	}
	lexer := New(input)
	for i, tt := range tests {
		tok := lexer.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestToken2(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	let add = fn(x, y) {
		x + y;
	};
	let result = add(five, ten);
	`
	tests := []struct {
		expectedType    tokens.Token
		expectedLiteral string
	}{
		{tokens.LET, "let"},
		{tokens.INDENTIFIER, "five"},
		{tokens.ASSGIN, "="},
		{tokens.INT, "5"},
		{tokens.SEMICOLON, ";"},
		{tokens.LET, "let"},
		{tokens.INDENTIFIER, "ten"},
		{tokens.ASSGIN, "="},
		{tokens.INT, "10"},
		{tokens.SEMICOLON, ";"},
		{tokens.LET, "let"},
		{tokens.INDENTIFIER, "add"},
		{tokens.ASSGIN, "="},
		{tokens.FUNCTION, "fn"},
		{tokens.LPARENT, "("},
		{tokens.INDENTIFIER, "x"},
		{tokens.COMMA, ","},
		{tokens.INDENTIFIER, "y"},
		{tokens.RPARENT, ")"},
		{tokens.LSQUIGL, "{"},
		{tokens.INDENTIFIER, "x"},
		{tokens.EQUAL, "+"},
		{tokens.INDENTIFIER, "y"},
		{tokens.SEMICOLON, ";"},
		{tokens.RSQUIGL, "}"},
		{tokens.SEMICOLON, ";"},

		{tokens.LET, "let"},
		{tokens.INDENTIFIER, "result"},
		{tokens.ASSGIN, "="},
		{tokens.INDENTIFIER, "add"},
		{tokens.LPARENT, "("},
		{tokens.INDENTIFIER, "five"},
		{tokens.COMMA, ","},
		{tokens.INDENTIFIER, "ten"},
		{tokens.RPARENT, ")"},
		{tokens.SEMICOLON, ";"},

		{tokens.EOF, ""},
	}
	lexer := New(input)
	for i, tt := range tests {
		tok := lexer.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
