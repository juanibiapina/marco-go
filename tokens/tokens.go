package tokens

type TokenType int

type Token struct {
	Typ   TokenType
	Value string
}

const (
	NUMBER TokenType = iota
)

func New(typ TokenType, value string) Token {
	return Token{typ, value}
}
