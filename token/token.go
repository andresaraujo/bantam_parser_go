package token

type TokenType string

const (
	// single-character tokens
	Plus  TokenType = "plus"
	Minus TokenType = "minus"
	Star  TokenType = "star"
	Slash TokenType = "slash"

	// literals
	Number     TokenType = "number"
	Str        TokenType = "string"
	Identifier TokenType = "identifier"

	TokenError TokenType = "error token"
	Eof        TokenType = "eof"
)

type Token struct {
	TokenType TokenType
	Lexeme    string
	Line      int
}
