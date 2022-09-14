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

	return &IdentifierExpression{name: t.Lexeme}, nil
}

type PrefixOperatorParselet struct {
	precedence Precedence
}

func (p *PrefixOperatorParselet) parse(parser *BantamParser, token token.Token) (Expression, error) {
	right, err := parser.Parse(p.precedence)

	if err != nil {
		return nil, err
	}

	return &PrefixExpression{operator: token, right: right}, nil
}

type GroupParselet struct{}

func (p *GroupParselet) parse(parser *BantamParser, t token.Token) (Expression, error) {
	expression, err := parser.Parse(none)
	if err != nil {
		return nil, err
	}

	if err := parser.consume(token.RightParen, fmt.Sprintf("Expect ')' after expression at line: %v", t.Line)); err != nil {
		return nil, err
	}
	return expression, nil
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

type ConditionParselet struct{}

func (p *ConditionParselet) getPrecedence() Precedence {
	return comparison
}

func (p *ConditionParselet) parse(parser *BantamParser, left Expression, t token.Token) (Expression, error) {
	thenExpression, err := parser.Parse(none)

	if err != nil {
		return nil, err
	}

	if err := parser.consume(token.Colon, fmt.Sprintf("Expected : after ? expression at %v", parser.current.Line)); err != nil {
		return nil, err
	}

	// Using one less precedence to make it right associative in case of multiple conditions
	// a ? b ? c : d : e -> (a ? (b ? c : d) : e)
	elseExpression, err := parser.Parse(p.getPrecedence() - 1)

	if err != nil {
		return nil, err
	}

	return &ConditionExpression{condition: left, thenExpression: thenExpression, elseExpression: elseExpression}, nil

}

type CallParselet struct{}

func (p *CallParselet) getPrecedence() Precedence {
	return call
}

func (p *CallParselet) parse(parser *BantamParser, left Expression, t token.Token) (Expression, error) {
	var arguments []Expression
	if parser.current.TokenType != token.RightParen {
		for {
			argument, err := parser.Parse(none)
			if err != nil {
				return nil, err
			}
			arguments = append(arguments, argument)
			if parser.current.TokenType != token.RightParen {
				if err := parser.consume(token.Comma, fmt.Sprintf("Expected ',' after argument at line: %v", parser.current.Line)); err != nil {
					return nil, err
				}
			} else {
				break
			}
		}
	}

	if err := parser.consume(token.RightParen, fmt.Sprintf("Expect ')' after arguments at line: %v", t.Line)); err != nil {
		return nil, err
	}

	return &CallExpression{callee: left, arguments: arguments}, nil
}

type AssignParselet struct{}

func (p *AssignParselet) getPrecedence() Precedence {
	return assignment
}

func (p *AssignParselet) parse(parser *BantamParser, left Expression, t token.Token) (Expression, error) {
	right, err := parser.Parse(p.getPrecedence() - 1)

	if err != nil {
		return nil, err
	}

	// if left is not an identifier, it is not a valid assignment
	if _, ok := left.(*IdentifierExpression); !ok {
		return nil, fmt.Errorf("left hand of an assignment must be an identifier, at line: %v", t.Line)
	}

	return &AssignmentExpression{identifier: left, operator: t, right: right}, nil
}

type PostfixOperatorParselet struct {
	precedence Precedence
}

func (p *PostfixOperatorParselet) getPrecedence() Precedence {
	return p.precedence
}

func (p *PostfixOperatorParselet) parse(parser *BantamParser, left Expression, t token.Token) (Expression, error) {
	return &PostfixExpression{left: left, operator: t}, nil
}
