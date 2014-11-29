package marco

import (
	"fmt"
	"strconv"
)

type parser struct {
	tokens chan Token

	currentToken Token
}

func (p *parser) next() {
	p.currentToken = <-p.tokens
}

func (p *parser) init() {
	p.next()
}

func (p *parser) errorf(format string, args ...interface{}) Expr {
	return Error{fmt.Sprintf(format, args...)}
}

func (p *parser) parseNumber() Expr {
	v, err := strconv.ParseInt(p.currentToken.value, 10, 64)
	if err != nil {
		return p.errorf("Error parsing number '%v': %s", p.currentToken.value, err)
	}
	return Number{v}
}

func Parse(tokens chan Token) Expr {
	p := &parser{
		tokens: tokens,
	}

	p.init()

	return p.parseNumber()
}
