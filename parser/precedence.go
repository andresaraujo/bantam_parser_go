package parser

type Precedence int

const (
	// precedence
	none Precedence = iota
	assignment
	equality
	comparison
	term
	product
	exponent
	unary
	postfix
	call
	primary
)
