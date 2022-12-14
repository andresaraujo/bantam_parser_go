package parser

import (
	"bantam_parser/lexer"
	"bantam_parser/token"
	"fmt"
	"log"
)

type BantamParser struct {
	lexer           *lexer.Lexer
	current         token.Token
	previous        token.Token
	hadError        bool
	prefixParselets map[token.TokenType]PrefixParselet
	infixParselets  map[token.TokenType]InfixParselet
}

func NewBantamParser(input string) *BantamParser {
	lexer := lexer.NewLexer(input)

	parser := &BantamParser{
		lexer:           lexer,
		current:         token.Token{TokenType: token.Eof, Lexeme: "", Line: 1},
		previous:        token.Token{TokenType: token.Eof, Lexeme: "", Line: 1},
		hadError:        false,
		prefixParselets: map[token.TokenType]PrefixParselet{},
		infixParselets:  map[token.TokenType]InfixParselet{},
	}

	parser.registerPrefix(token.Identifier, &PrimaryParselet{})
	parser.registerPrefix(token.Number, &PrimaryParselet{})
	parser.registerPrefix(token.LeftParen, &GroupParselet{})

	parser.registerInfix(token.Equal, &AssignParselet{})
	parser.registerInfix(token.LeftParen, &CallParselet{})
	parser.registerInfix(token.Question, &ConditionParselet{})
	parser.registerInfix(token.Bang, &PostfixOperatorParselet{precedence: postfix})

	parser.registerPrefix(token.Plus, &PrefixOperatorParselet{precedence: unary})
	parser.registerPrefix(token.Minus, &PrefixOperatorParselet{precedence: unary})
	parser.registerPrefix(token.Bang, &PrefixOperatorParselet{precedence: unary})
	parser.registerPrefix(token.Tilde, &PrefixOperatorParselet{precedence: unary})

	parser.registerInfix(token.Plus, &BinaryOperatorParselet{precedence: term, associativity: left})
	parser.registerInfix(token.Minus, &BinaryOperatorParselet{precedence: term, associativity: left})
	parser.registerInfix(token.Star, &BinaryOperatorParselet{precedence: product, associativity: left})
	parser.registerInfix(token.Slash, &BinaryOperatorParselet{precedence: product, associativity: left})
	parser.registerInfix(token.Caret, &BinaryOperatorParselet{precedence: exponent, associativity: right})

	return parser
}

func (p *BantamParser) isAtEnd() bool {
	return p.current.TokenType == token.Eof
}

func (p *BantamParser) checkType(tokenType token.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return tokenType == p.current.TokenType
}

func (p *BantamParser) consume(tokenType token.TokenType, errorMessage string) error {
	if p.checkType(tokenType) {
		p.Advance()
		return nil
	}
	return fmt.Errorf("%v", errorMessage)
}

func (p *BantamParser) Advance() {
	p.previous = p.current

	for {
		current, err := p.lexer.ScanToken()
		p.current = current

		if p.current.TokenType != token.TokenError {
			break
		}
		if err == nil {
			p.hadError = true
			log.Fatal(err)
			return
		}
	}
}

func (p *BantamParser) Parse(precedence Precedence) (Expression, error) {
	p.Advance()

	prefix := p.prefixParselets[p.previous.TokenType]
	if prefix == nil {
		return nil, fmt.Errorf("no prefix parselet for %v", p.previous.TokenType)
	}
	left, err := prefix.parse(p, p.previous)
	if err != nil {
		return nil, err
	}

	for precedence < p.getPrecedence() {

		p.Advance()
		infix := p.infixParselets[p.previous.TokenType]
		if infix == nil {
			return nil, fmt.Errorf("no infix parselet for %v", p.previous.TokenType)
		}
		left, err = infix.parse(p, left, p.previous)
		if err != nil {
			return nil, err
		}
	}

	return left, nil
}

func (p *BantamParser) registerPrefix(tokenType token.TokenType, parselet PrefixParselet) {
	p.prefixParselets[tokenType] = parselet
}

func (p *BantamParser) registerInfix(tokenType token.TokenType, parselet InfixParselet) {
	p.infixParselets[tokenType] = parselet
}

func (p *BantamParser) getPrecedence() Precedence {
	infix := p.infixParselets[p.current.TokenType]
	if infix == nil {
		return 0
	}

	return infix.getPrecedence()
}
