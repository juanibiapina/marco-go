package scanner

import (
	"github.com/juanibiapina/marco/tokens"
	"testing"
	"time"
)

func validateToken(t *testing.T, token tokens.Token, typ tokens.TokenType, value string) {
	if token.Typ != typ {
		t.Errorf("Wrong token type in '%v': expected '%v', got '%v'", token.Value, typ, token.Typ)
	}
	if token.Value != value {
		t.Errorf("Wrong token value: expected '%v', got '%v'", value, token.Value)
	}
}

func assertNextToken(t *testing.T, c chan tokens.Token, typ tokens.TokenType, value string) {
	select {
	case token := <-c:
		validateToken(t, token, typ, value)
	case <-time.After(2 * time.Second):
		t.Errorf("Expected '%v' but did not produce a token in time", value)
	}
}

var tokenTests = []struct {
	input string
	typ   tokens.TokenType
	value string
}{
	{"1", tokens.NUMBER, "1"},
	{"2", tokens.NUMBER, "2"},
	{"823", tokens.NUMBER, "823"},
	{"def", tokens.NAME, "def"},
	{":asdf", tokens.SYMBOL, "asdf"},
	{"\"\"", tokens.STRING, ""},
	{"", tokens.EOF, ""},
}

func TestScanTokens(t *testing.T) {
	for _, tt := range tokenTests {
		c := Scan([]byte(tt.input))
		assertNextToken(t, c, tt.typ, tt.value)
	}
}

func TestScanList(t *testing.T) {
	c := Scan([]byte("[1 2 3]"))

	assertNextToken(t, c, tokens.LBRACKET, "[")
	assertNextToken(t, c, tokens.NUMBER, "1")
	assertNextToken(t, c, tokens.NUMBER, "2")
	assertNextToken(t, c, tokens.NUMBER, "3")
	assertNextToken(t, c, tokens.RBRACKET, "]")
}

func TestScanMultiLine(t *testing.T) {
	c := Scan([]byte("1\n\n2"))

	assertNextToken(t, c, tokens.NUMBER, "1")
	assertNextToken(t, c, tokens.NUMBER, "2")
}

func TestScanApplication(t *testing.T) {
	c := Scan([]byte("(def :a 1)"))

	assertNextToken(t, c, tokens.LPAREN, "(")
	assertNextToken(t, c, tokens.NAME, "def")
	assertNextToken(t, c, tokens.SYMBOL, "a")
	assertNextToken(t, c, tokens.NUMBER, "1")
	assertNextToken(t, c, tokens.RPAREN, ")")
}
