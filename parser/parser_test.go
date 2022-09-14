package parser

import (
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	// Function call.
	expect(t, "a()", "a()")
	expect(t, "a(b)", "a(b)")
	expect(t, "a(b, c)", "a(b, c)")
	expect(t, "a(b)(c)", "a(b)(c)")
	expect(t, "a(b) + c(d)", "(a(b) + c(d))")
	expect(t, "a(b ? c : d, e + f)", "a((b ? c : d), (e + f))")

	// Unary precedence.
	expect(t, "~!-+a", "(~(!(-(+a))))")
	expect(t, "a!!!", "(((a!)!)!)")

	// Unary and binary precedence.
	expect(t, "-a * b", "((-a) * b)")
	expect(t, "!a + b", "((!a) + b)")
	expect(t, "~a ^ b", "((~a) ^ b)")
	expect(t, "-a!", "(-(a!))")
	expect(t, "!a!", "(!(a!))")

	// Binary precedence.
	expect(t, "a = b + c * d ^ e - f / g", "(a = ((b + (c * (d ^ e))) - (f / g)))")

	// Binary associativity.
	expect(t, "a = b = c", "(a = (b = c))")
	expect(t, "a + b - c", "((a + b) - c)")
	expect(t, "a * b / c", "((a * b) / c)")
	expect(t, "a ^ b ^ c", "(a ^ (b ^ c))")

	// Conditional operator.
	expect(t, "a ? b : c ? d : e", "(a ? b : (c ? d : e))")
	expect(t, "a ? b ? c : d : e", "(a ? (b ? c : d) : e)")
	expect(t, "a + b ? c * d : e / f", "((a + b) ? (c * d) : (e / f))")

	// Grouping.
	expect(t, "a + (b + c) + d", "((a + (b + c)) + d)")
	expect(t, "a ^ (b + c)", "(a ^ (b + c))")
	expect(t, "(!a)!", "((!a)!)")
}

func expect(t *testing.T, source string, expected string) {
	parser := NewBantamParser(source)
	parser.Advance()

	expression, err := parser.Parse(0)

	if err != nil {
		t.Error(err)
	}

	strBuilder := &strings.Builder{}
	expression.Print(strBuilder)
	result := strBuilder.String()

	if expected != result {
		t.Error("Expected", expected, "got", result)
	}
}
