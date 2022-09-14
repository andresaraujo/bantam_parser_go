package token

type TokenType string

const (
	// single-character tokens
	Plus       TokenType = "plus"        // +
	Minus      TokenType = "minus"       // -
	Star       TokenType = "star"        // *
	Slash      TokenType = "slash"       // /
	LeftParen  TokenType = "left_paren"  // (
	RightParen TokenType = "right_paren" // )
	Caret      TokenType = "caret"       // ^
	Bang       TokenType = "bang"        // !
	Tilde      TokenType = "tilde"       // ~
	Equal      TokenType = "equal"       // =
	Question   TokenType = "question"    // ?
	Colon      TokenType = "colon"       // :
	Comma      TokenType = "comma"       // ,

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
