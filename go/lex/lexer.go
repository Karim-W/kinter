package lex

import "github.com/karim-w/kinter/go/tokens"

type Entity struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Entity {
	l := &Entity{input: input}
	l.readChar()
	return l
}

func (l *Entity) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for "NUL" (null character)
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Entity) NextToken() tokens.Entity {
	var tok tokens.Entity

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(tokens.ASSGIN, l.ch)
	case '+':
		tok = newToken(tokens.EQUAL, l.ch)
	case '(':
		tok = newToken(tokens.LPARENT, l.ch)
	case ')':
		tok = newToken(tokens.RPARENT, l.ch)
	case '{':
		tok = newToken(tokens.LSQUIGL, l.ch)
	case '}':
		tok = newToken(tokens.RSQUIGL, l.ch)
	case ',':
		tok = newToken(tokens.COMMA, l.ch)
	case ';':
		tok = newToken(tokens.SEMICOLON, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = tokens.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = tokens.LookupIdent(tok.Literal)
			return tok
		}
		if isDigit(l.ch) {
			tok.Type = tokens.INT
			tok.Literal = l.readNumber()
			return tok
		}
		tok = newToken(tokens.ILLEGAL, l.ch)
	}
	l.readChar()
	return tok
}

func newToken(tokenType tokens.Token, ch byte) tokens.Entity {
	return tokens.Entity{Type: tokenType, Literal: string(ch)}
}

func (l *Entity) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Entity) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func (l *Entity) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
