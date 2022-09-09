package parser

type Precedence int

const (
	// precedence
	none Precedence = iota
	assignment
	term
	product
	exponent
	unary
	call
	primary
)
