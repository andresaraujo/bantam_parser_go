package parser

import (
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	expect(t, "1 + 2", "(1 + 2)")
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
