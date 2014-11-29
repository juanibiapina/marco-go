package parser

import (
	"github.com/juanibiapina/marco/lang"
	"github.com/juanibiapina/marco/tokens"
	"log"
	"strconv"
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

func (p *parser) parseForm() lang.Expr {
	switch p.currentToken.Typ {
	case tokens.NUMBER:
		return p.parseNumber()
	case tokens.NAME:
		return p.parseName()
	default:
		p.errorf("Parse error: unexpected token '%v'", p.currentToken)
		return nil
	}
}

func (p *parser) parseName() lang.Expr {
	return lang.Name{p.currentToken.Value}
}

func (p *parser) parseNumber() lang.Expr {
	v, err := strconv.ParseInt(p.currentToken.Value, 10, 64)
	if err != nil {
		p.errorf("Error parsing number '%v': %s", p.currentToken.Value, err)
	}
	return lang.Number{v}
}

func Parse(tokens chan tokens.Token) lang.Expr {
	p := &parser{
		tokens: tokens,
	}

	p.init()

	return p.parseForm()
}
