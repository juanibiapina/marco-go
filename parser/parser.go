package parser

import (
	"github.com/juanibiapina/marco/tokens"
	"log"
)

type parser struct {
	tokens chan tokens.Token

	currentToken tokens.Token
}

func (p *parser) next() {
	p.currentToken = <-p.tokens
}

func (p *parser) init() {
	p.next()
}

func (p *parser) errorf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}
