package scanner

import (
	"github.com/juanibiapina/marco/tokens"
	"log"
	"strings"
	"unicode/utf8"
)

type scanner struct {
	input  []byte
	tokens chan tokens.Token
	start  int
	pos    int
	width  int
}

func (l *scanner) next() (r rune) {
	r, l.width = utf8.DecodeRune(l.input[l.pos:])
	if l.width == 0 {
		return -1
	}
	l.pos += l.width
	return r
}

func (l *scanner) ignore() {
	l.start = l.pos
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

func (l *scanner) acceptRunFunc(f func(rune) bool) {
	for f(l.next()) {
	}
	l.backup()
}

func (l *scanner) errorf(format string, args ...interface{}) {
	log.Fatalf(format, args) // print line and column information
}

func (l *scanner) run() {
	for state := lexInitial; state != nil; {
		state = state(l)
	}
	close(l.tokens)
}
