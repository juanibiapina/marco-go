package parser

import (
	"github.com/juanibiapina/marco/lang"
	"github.com/juanibiapina/marco/tokens"
)

func Parse(tokens chan tokens.Token) lang.Block {
	p := &parser{
		tokens: tokens,
	}

	p.init()

	return p.parseTopLevel()
}
