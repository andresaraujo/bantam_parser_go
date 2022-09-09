package parser

import (
	"bantam_parser/token"
	"fmt"
)

type PrefixParselet interface {
	parse(parser *BantamParser, token token.Token) (Expression, error)
}

type InfixParselet interface {
	parse(parser *BantamParser, left Expression, token token.Token) (Expression, error)
	getPrecedence() Precedence
}

type PrimaryParselet struct{}

func (p PrimaryParselet) parse(parser *BantamParser, t token.Token) (Expression, error) {
	switch t.TokenType {
	case token.Number:
		return LiteralExpression{value: t.Lexeme}, nil
	}
	return nil, fmt.Errorf("invalid token %v", t.TokenType)
}

type PrefixOperatorParselet struct {
	precedence Precedence
}

func (p *PrefixOperatorParselet) parse(parser *BantamParser, token token.Token) (Expression, error) {
	right, err := parser.Parse(p.precedence)
	return &PrefixExpression{operator: token, right: right}, err
}

type BinaryOperatorParselet struct {
	precedence    Precedence
	associativity Associativity
}

func (p *BinaryOperatorParselet) getPrecedence() Precedence {
	return p.precedence
}

func (p *BinaryOperatorParselet) parse(parser *BantamParser, left Expression, token token.Token) (Expression, error) {
	precedence := p.precedence
	if p.associativity == right {
		precedence--
	}
	right, err := parser.Parse(precedence)
	return &BinaryExpression{left: left, operator: token, right: right}, err
}
