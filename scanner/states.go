package scanner

import (
	"github.com/juanibiapina/marco/tokens"
	"unicode"
)

type stateFn func(*scanner) stateFn

func scanNumber(l *scanner) stateFn {
	l.acceptRun("0123456789")
	l.emit(tokens.NUMBER)

	return scanInitial
}

func scanName(l *scanner) stateFn {
	l.acceptLexeme(lexIdentifier)
	l.emit(tokens.NAME)

	return scanInitial
}

func scanSymbol(l *scanner) stateFn {
	l.ignore() // ignore the ':' TODO panic on wrong ignore
	l.acceptLexeme(lexIdentifier)
	l.emit(tokens.SYMBOL)

	return scanInitial
}

func scanString(l *scanner) stateFn {
	l.ignore()
	l.acceptUntil('"')
	l.emit(tokens.STRING)
	l.accept('"')
	l.ignore()

	return scanInitial
}

func scanInitial(l *scanner) stateFn {
	r := l.next()

	if r == -1 {
		l.emit(tokens.EOF)
		return nil
	}

	if r == ':' {
		return scanSymbol
	}

	if r == '"' {
		return scanString
	}

	if r == '.' {
		l.emit(tokens.DOT)
		return scanInitial
	}

	if unicode.IsDigit(r) {
		l.backup()
		return scanNumber
	}

	if unicode.IsLetter(r) {
		l.backup()
		return scanName
	}

	if unicode.IsSpace(r) {
		l.ignore()
		return scanInitial
	}

	if r == '(' {
		l.emit(tokens.LPAREN)
		return scanInitial
	}

	if r == ')' {
		l.emit(tokens.RPAREN)
		return scanInitial
	}

	if r == '[' {
		l.emit(tokens.LBRACKET)
		return scanInitial
	}

	if r == ']' {
		l.emit(tokens.RBRACKET)
		return scanInitial
	}

	l.errorf("Unrecognized character: %v", string(r))
	return nil
}
