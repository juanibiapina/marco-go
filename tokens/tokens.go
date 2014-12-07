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
	SYMBOL
	STRING
	DOT
	EOF
	LPAREN
	RPAREN
	LBRACKET
	RBRACKET
)

func New(typ TokenType, value string) Token {
	return Token{typ, value}
}

var typeNames = []string{
	"TOKEN_NONE",
	"TOKEN_NUMBER",
	"TOKEN_NAME",
	"TOKEN_SYMBOL",
	"TOKEN_STRING",
	"TOKEN_DOT",
	"TOKEN_EOF",
	"TOKEN_LPAREN",
	"TOKEN_RPAREN",
	"TOKEN_LBRACKET",
	"TOKEN_RBRACKET",
}

func (typ TokenType) String() string {
	return typeNames[int(typ)]
}
