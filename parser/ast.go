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
