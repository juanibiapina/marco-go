package scanner

import (
	"github.com/juanibiapina/marco/tokens"
	"strings"
	"unicode/utf8"
)

type scanner struct {
	input  string
	tokens chan tokens.Token
	start  int
	pos    int
	width  int
}

type stateFn func(*scanner) stateFn

func (l *scanner) next() (r rune) {
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}

func (l *scanner) backup() {
	l.pos -= l.width
}

func (l *scanner) emit(typ tokens.TokenType) {
	l.tokens <- tokens.New(typ, l.input[l.start:l.pos])
	l.start = l.pos
}

func (l *scanner) acceptRun(values string) {
	for strings.IndexRune(values, l.next()) >= 0 {
	}
	l.backup()
}

func numberState(l *scanner) stateFn {
	l.acceptRun("0123456789")
	l.emit(tokens.NUMBER)

	l.tokens <- tokens.New(tokens.NUMBER, "2")
	return nil
}

var initialState = numberState

func (l *scanner) run() {
	for state := initialState; state != nil; {
		state = state(l)
	}
}

func Scan(input string) chan tokens.Token {
	scanner := &scanner{
		input:  input,
		tokens: make(chan tokens.Token),
	}
	go scanner.run()
	return scanner.tokens
}
