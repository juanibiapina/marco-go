package parser

import (
	"fmt"
	"github.com/juanibiapina/marco/lang"
	"github.com/juanibiapina/marco/tokens"
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

func (p *parser) errorf(format string, args ...interface{}) lang.Expr {
	return lang.Error{fmt.Sprintf(format, args...)}
}

func (p *parser) parseNumber() lang.Expr {
	v, err := strconv.ParseInt(p.currentToken.Value, 10, 64)
	if err != nil {
		return p.errorf("Error parsing number '%v': %s", p.currentToken.Value, err)
	}
	return lang.Number{v}
}

func Parse(tokens chan tokens.Token) lang.Expr {
	p := &parser{
		tokens: tokens,
	}

	p.init()

	return p.parseNumber()
}
