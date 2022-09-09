package lexer

import (
	"bantam_parser/token"
	"testing"
)

func TestScanner(t *testing.T) {
	lexer := NewLexer("1 + 2")

	expectToken(t, lexer, token.Number)
	expectToken(t, lexer, token.Plus)
	expectToken(t, lexer, token.Number)
}

func expectToken(t *testing.T, lexer *Lexer, expectedTokenType token.TokenType) {
	token, err := lexer.ScanToken()

	if err != nil {
		t.Error(err)
	}

	if token.TokenType != expectedTokenType {
		t.Error("Expected token type", expectedTokenType, "got", token.TokenType)
	}
}
