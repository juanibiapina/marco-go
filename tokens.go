package marco

type TokenType int

type Token struct {
	typ   TokenType
	value string
}

const (
	TOKEN_NUMBER TokenType = iota
)
