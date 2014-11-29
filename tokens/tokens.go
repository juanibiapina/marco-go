package tokens

type TokenType int

type Token struct {
	Typ   TokenType
	Value string
}

const (
	NUMBER TokenType = iota
	NAME
)

func New(typ TokenType, value string) Token {
	return Token{typ, value}
}

var typeNames = []string{
	"TOKEN_NUMBER",
	"TOKEN_NAME",
}

func (typ TokenType) String() string {
	return typeNames[int(typ)]
}
