package scanner

import (
	"github.com/juanibiapina/marco/tokens"
)

func Scan(input []byte) chan tokens.Token {
	scanner := &scanner{
		input:  input,
		tokens: make(chan tokens.Token),
	}
	go scanner.run()
	return scanner.tokens
}
