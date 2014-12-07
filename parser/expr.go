package parser

import (
	"github.com/juanibiapina/marco/lang"
	"github.com/juanibiapina/marco/tokens"
	"strconv"
)

func (p *parser) parseTopLevel() lang.Expr {
	var forms []lang.Expr

	for p.currentToken.Typ != tokens.EOF {
		if p.currentToken.Typ == tokens.COMMENT {
			p.next()
			continue
		}

		forms = append(forms, p.parseForm())
	}

	return lang.MakeBlock(lang.SliceToList(forms))
}

func (p *parser) parseForm() lang.Expr {
	switch p.currentToken.Typ {
	case tokens.NUMBER:
		return p.parseNumber()
	case tokens.SYMBOL:
		return p.parseSymbol()
	case tokens.NAME:
		return p.parseName()
	case tokens.STRING:
		return p.parseString()
	case tokens.LBRACKET:
		p.next()
		return p.parseList()
	case tokens.LPAREN:
		p.next()
		return p.parseApplication()
	default:
		p.errorf("Parse error: unexpected token '%v'", p.currentToken)
		return nil
	}
}

func (p *parser) parseApplication() lang.Expr {
	var list []lang.Expr

	for p.currentToken.Typ != tokens.RPAREN {
		expr := p.parseForm()
		list = append(list, expr)
	}
	p.accept(tokens.RPAREN)
	return lang.MakeApplication(lang.SliceToList(list))
}

func (p *parser) parseList() lang.Expr {
	var list []lang.Expr

	for p.currentToken.Typ != tokens.RBRACKET {
		expr := p.parseForm()
		list = append(list, expr)
	}
	p.accept(tokens.RBRACKET)
	return lang.SliceToList(list)
}

func (p *parser) parseName() lang.Expr {
	value := p.currentToken.Value
	p.next()
	if p.currentToken.Typ == tokens.DOT {
		p.next()
		return lang.MakeNestedName(value, p.parseName())
	} else {
		return lang.MakeName(value)
	}
}

func (p *parser) parseString() lang.Expr {
	result := lang.MakeString(p.currentToken.Value)
	p.next()
	return result
}

func (p *parser) parseSymbol() lang.Expr {
	result := lang.MakeSymbol(p.currentToken.Value)
	p.next()
	return result
}

func (p *parser) parseNumber() lang.Expr {
	v, err := strconv.ParseInt(p.currentToken.Value, 10, 64)
	if err != nil {
		p.errorf("Error parsing number '%v': %s", p.currentToken.Value, err)
	}
	p.next()
	return lang.MakeNumber(v)
}
