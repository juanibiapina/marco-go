package scanner

import (
	"github.com/juanibiapina/marco/tokens"
	"unicode"
)

type stateFn func(*scanner) stateFn

func lexNumber(l *scanner) stateFn {
	l.acceptRun("0123456789")
	l.emit(tokens.NUMBER)

	return nil
}

func lexName(l *scanner) stateFn {
	l.acceptRunFunc(unicode.IsLetter)
	l.emit(tokens.NAME)

	return nil
}

func lexForm(l *scanner) stateFn {
	r := l.next()

	if unicode.IsDigit(r) {
		l.backup()
		return lexNumber
	}

	if unicode.IsLetter(r) {
		l.backup()
		return lexName
	}

	l.errorf("Unrecognized character: %v", string(r))
	return nil
}

var lexInitial = lexForm
