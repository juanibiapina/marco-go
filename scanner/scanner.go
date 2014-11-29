package scanner

import (
	"github.com/juanibiapina/marco/tokens"
	"log"
	"strings"
	"unicode"
	"unicode/utf8"
)

type scanner struct {
	input  []byte
	tokens chan tokens.Token
	start  int
	pos    int
	width  int
}

type stateFn func(*scanner) stateFn

func (l *scanner) next() (r rune) {
	r, l.width = utf8.DecodeRune(l.input[l.pos:])
	l.pos += l.width
	return r
}

func (l *scanner) backup() {
	l.pos -= l.width
}

func (l *scanner) emit(typ tokens.TokenType) {
	l.tokens <- tokens.New(typ, string(l.input[l.start:l.pos]))
	l.start = l.pos
}

func (l *scanner) acceptRun(values string) {
	for strings.IndexRune(values, l.next()) >= 0 {
	}
	l.backup()
}

func lexNumber(l *scanner) stateFn {
	l.acceptRun("0123456789")
	l.emit(tokens.NUMBER)

	return nil
}

func (l *scanner) acceptRunFunc(f func(rune) bool) {
	for f(l.next()) {
	}
	l.backup()
}

func lexName(l *scanner) stateFn {
	l.acceptRunFunc(unicode.IsLetter)
	l.emit(tokens.NAME)

	return nil
}

func (l *scanner) errorf(format string, args ...interface{}) {
	log.Fatalf(format, args) // print line and column information
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

func (l *scanner) run() {
	for state := lexInitial; state != nil; {
		state = state(l)
	}
}

func Scan(input []byte) chan tokens.Token {
	scanner := &scanner{
		input:  input,
		tokens: make(chan tokens.Token),
	}
	go scanner.run()
	return scanner.tokens
}
