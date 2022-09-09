package ascii

// ascii codes enum
const (
	LeftParen  byte = 0x28
	RightParen byte = 0x29
	Comma      byte = 0x2C
	Minus      byte = 0x2D
	Plus       byte = 0x2B
	Slash      byte = 0x2F
	Star       byte = 0x2A
	Num0       byte = 0x30
	Num9       byte = 0x39
	A          byte = 0x41
	Z          byte = 0x5A
	a          byte = 0x61
	z          byte = 0x7A
)

func IsDigit(charcode byte) bool {
	return charcode >= Num0 && charcode <= Num9
}
