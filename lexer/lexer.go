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

	if ascii.IsAlpha(char) {
		return l.scanIdentifier()
	}

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
	case ascii.Comma:
		return l.makeToken(token.Comma), nil
	case ascii.Colon:
		return l.makeToken(token.Colon), nil
	case ascii.Caret:
		return l.makeToken(token.Caret), nil
	case ascii.Bang:
		return l.makeToken(token.Bang), nil
	case ascii.LeftParen:
		return l.makeToken(token.LeftParen), nil
	case ascii.RightParen:
		return l.makeToken(token.RightParen), nil
	case ascii.Equal:
		return l.makeToken(token.Equal), nil
	case ascii.Tilde:
		return l.makeToken(token.Tilde), nil
	case ascii.Question:
		return l.makeToken(token.Question), nil
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

func (l *Lexer) scanIdentifier() (token.Token, error) {
	for ascii.IsAlpha(l.peek()) || ascii.IsDigit(l.peek()) {
		l.advance()
	}

	text := l.input[l.start:l.current]

	return token.Token{TokenType: token.Identifier, Lexeme: text, Line: l.line}, nil
}
