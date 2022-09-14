package parser

import (
	"bantam_parser/token"
	"fmt"
	"strings"
)

type Expression interface {
	Print(strBuilder *strings.Builder)
}
type LiteralExpression struct {
	value interface{}
}

func (e LiteralExpression) Print(strBuilder *strings.Builder) {
	strBuilder.WriteString(fmt.Sprintf("%v", e.value))
}

type PrefixExpression struct {
	operator token.Token
	right    Expression
}

func (e PrefixExpression) Print(strBuilder *strings.Builder) {
	strBuilder.WriteString("(")
	strBuilder.WriteString(e.operator.Lexeme)
	e.right.Print(strBuilder)
	strBuilder.WriteString(")")
}

type BinaryExpression struct {
	left     Expression
	operator token.Token
	right    Expression
}

func (e *BinaryExpression) Print(strBuilder *strings.Builder) {
	strBuilder.WriteString("(")
	e.left.Print(strBuilder)
	strBuilder.WriteString(" " + e.operator.Lexeme + " ")
	e.right.Print(strBuilder)
	strBuilder.WriteString(")")
}

type ConditionExpression struct {
	condition      Expression
	thenExpression Expression
	elseExpression Expression
}

func (e *ConditionExpression) Print(strBuilder *strings.Builder) {
	strBuilder.WriteString("(")
	e.condition.Print(strBuilder)
	strBuilder.WriteString(" ? ")
	e.thenExpression.Print(strBuilder)
	strBuilder.WriteString(" : ")
	e.elseExpression.Print(strBuilder)
	strBuilder.WriteString(")")
}

type CallExpression struct {
	callee    Expression
	arguments []Expression
}

func (e *CallExpression) Print(strBuilder *strings.Builder) {
	e.callee.Print(strBuilder)
	strBuilder.WriteString("(")
	for i, arg := range e.arguments {
		if i > 0 {
			strBuilder.WriteString(", ")
		}
		arg.Print(strBuilder)
	}
	strBuilder.WriteString(")")
}

type AssignmentExpression struct {
	identifier Expression
	operator   token.Token
	right      Expression
}

func (e *AssignmentExpression) Print(strBuilder *strings.Builder) {
	strBuilder.WriteString("(")
	e.identifier.Print(strBuilder)
	strBuilder.WriteString(" " + e.operator.Lexeme + " ")
	e.right.Print(strBuilder)
	strBuilder.WriteString(")")
}

type IdentifierExpression struct {
	name string
}

func (e *IdentifierExpression) Print(strBuilder *strings.Builder) {
	strBuilder.WriteString(e.name)
}

type PostfixExpression struct {
	left     Expression
	operator token.Token
}

func (e *PostfixExpression) Print(strBuilder *strings.Builder) {
	strBuilder.WriteString("(")
	e.left.Print(strBuilder)
	strBuilder.WriteString(e.operator.Lexeme)
	strBuilder.WriteString(")")
}
