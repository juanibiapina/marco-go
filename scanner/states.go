package scanner

import (
	"github.com/juanibiapina/marco/tokens"
	"unicode"
)

type stateFn func(*scanner) stateFn

func lexNumber(l *scanner) stateFn {
	l.acceptRun("0123456789")
	l.emit(tokens.NUMBER)

	return lexInitial
}

func lexName(l *scanner) stateFn {
	l.acceptRunFunc(unicode.IsLetter)
	l.emit(tokens.NAME)

	return lexInitial
}

func lexSymbol(l *scanner) stateFn {
	l.ignore()
	l.acceptRunFunc(unicode.IsLetter)
	l.emit(tokens.SYMBOL)

	return lexInitial
}

func lexString(l *scanner) stateFn {
	l.ignore()
	l.acceptUntil('"')
	l.emit(tokens.STRING)
	l.accept('"')
	l.ignore()

	return lexInitial
}

func lexInitial(l *scanner) stateFn {
	r := l.next()

	if r == -1 {
		l.emit(tokens.EOF)
		return nil
	}

	if r == ':' {
		return lexSymbol
	}

	if r == '"' {
		return lexString
	}

	if r == '.' {
		l.emit(tokens.DOT)
		return lexInitial
	}

	if unicode.IsDigit(r) {
		l.backup()
		return lexNumber
	}

	if unicode.IsLetter(r) {
		l.backup()
		return lexName
	}

	if unicode.IsSpace(r) {
		l.ignore()
		return lexInitial
	}

	if r == '(' {
		l.emit(tokens.LPAREN)
		return lexInitial
	}

	if r == ')' {
		l.emit(tokens.RPAREN)
		return lexInitial
	}

	if r == '[' {
		l.emit(tokens.LBRACKET)
		return lexInitial
	}

	if r == ']' {
		l.emit(tokens.RBRACKET)
		return lexInitial
	}

	l.errorf("Unrecognized character: %v", string(r))
	return nil
}
