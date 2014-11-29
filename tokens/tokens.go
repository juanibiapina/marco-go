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
	LPAREN
	RPAREN
)

func New(typ TokenType, value string) Token {
	return Token{typ, value}
}

var typeNames = []string{
	"TOKEN_NONE",
	"TOKEN_NUMBER",
	"TOKEN_NAME",
	"TOKEN_EOF",
	"TOKEN_LPAREN",
	"TOKEN_RPAREN",
}

func (typ TokenType) String() string {
	return typeNames[int(typ)]
}
