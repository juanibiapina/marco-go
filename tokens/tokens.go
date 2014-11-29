package tokens

type TokenType int

type Token struct {
	Typ   TokenType
	Value string
}

const (
	NONE TokenType = iota
	NUMBER
	NAME
	EOF
)

func New(typ TokenType, value string) Token {
	return Token{typ, value}
}

var typeNames = []string{
	"TOKEN_NONE",
	"TOKEN_NUMBER",
	"TOKEN_NAME",
	"TOKEN_EOF",
}

func (typ TokenType) String() string {
	return typeNames[int(typ)]
}
