package parser

import (
	"github.com/juanibiapina/marco/lang"
	"github.com/juanibiapina/marco/tokens"
	"strconv"
)

func (p *parser) parseModule() lang.Expr {
	var forms []lang.Expr

	for p.currentToken.Typ != tokens.EOF {
		forms = append(forms, p.parseForm())
	}

	return lang.MakeModule(forms)
}

func (p *parser) parseForm() lang.Expr {
	switch p.currentToken.Typ {
	case tokens.NUMBER:
		return p.parseNumber()
	case tokens.NAME:
		return p.parseName()
	case tokens.LBRACKET:
		p.next()
		return p.parseList()
	default:
		p.errorf("Parse error: unexpected token '%v'", p.currentToken)
		return nil
	}
}

func (p *parser) parseList() lang.Expr {
	var list []lang.Expr

	for p.currentToken.Typ != tokens.RBRACKET {
		expr := p.parseForm()
		list = append(list, expr)
	}
	p.accept(tokens.RBRACKET)
	return lang.MakeList(list)
}

func (p *parser) parseName() lang.Expr {
	result := lang.Name{p.currentToken.Value}
	p.next()
	return result
}

func (p *parser) parseNumber() lang.Expr {
	v, err := strconv.ParseInt(p.currentToken.Value, 10, 64)
	if err != nil {
		p.errorf("Error parsing number '%v': %s", p.currentToken.Value, err)
	}
	p.next()
	return lang.Number{v}
}
