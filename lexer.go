package marco

import (
	"strings"
	"unicode/utf8"
)

type lexer struct {
	input  string
	tokens chan Token
	start  int
	pos    int
	width  int
}

type stateFn func(*lexer) stateFn

func (l *lexer) next() (r rune) {
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}

func (l *lexer) backup() {
	l.pos -= l.width
}

func (l *lexer) emit(typ TokenType) {
	l.tokens <- Token{typ, l.input[l.start:l.pos]}
	l.start = l.pos
}

func (l *lexer) acceptRun(values string) {
	for strings.IndexRune(values, l.next()) >= 0 {
	}
	l.backup()
}

func numberState(l *lexer) stateFn {
	l.acceptRun("0123456789")
	l.emit(TOKEN_NUMBER)

	l.tokens <- Token{TOKEN_NUMBER, "2"}
	return nil
}

var initialState = numberState

func (l *lexer) run() {
	for state := initialState; state != nil; {
		state = state(l)
	}
}

func Lex(input string) chan Token {
	lexer := &lexer{
		input:  input,
		tokens: make(chan Token),
	}
	go lexer.run()
	return lexer.tokens
}
