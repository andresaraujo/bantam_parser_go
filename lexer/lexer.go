package lexer

import (
	"bantam_parser/ascii"
	"bantam_parser/token"
)

type Lexer struct {
	input   string
	start   int
	current int
	line    int
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: input, start: 0, current: 0, line: 1}
}

func (l *Lexer) ScanToken() (token.Token, error) {
	l.skipWhitespace()

	l.start = l.current
	if l.isAtEnd() {
		return token.Token{TokenType: token.Eof, Lexeme: "", Line: l.line}, nil
	}

	char := l.nextChar()

	if ascii.IsDigit(char) {
		return l.number(), nil
	}

	switch char {
	case ascii.Plus:
		return l.makeToken(token.Plus), nil
	case ascii.Minus:
		return l.makeToken(token.Minus), nil
	case ascii.Star:
		return l.makeToken(token.Star), nil
	case ascii.Slash:
		return l.makeToken(token.Slash), nil
	}

	return token.Token{TokenType: token.TokenError, Lexeme: "", Line: l.line}, nil
}

func (l *Lexer) advance() {
	l.current++
}

func (l *Lexer) isAtEnd() bool {
	return l.current >= len(l.input)
}

func (l *Lexer) nextChar() byte {
	l.advance()
	return l.input[l.current-1]
}

func (l *Lexer) peek() byte {
	if l.current >= len(l.input) {
		return 0
	}
	return l.input[l.current]
}

func (l *Lexer) skipWhitespace() {
	for {
		c := l.peek()
		switch c {
		case ' ', '\r', '\t':
			l.advance()
		case '\n':
			l.line++
			l.advance()
		default:
			return
		}
	}
}

func (l *Lexer) makeToken(tokenType token.TokenType) token.Token {
	return token.Token{TokenType: tokenType, Lexeme: l.input[l.start:l.current], Line: l.line}
}

func (l *Lexer) number() token.Token {
	for ascii.IsDigit(l.peek()) {
		l.advance()
	}

	return l.makeToken(token.Number)
}
