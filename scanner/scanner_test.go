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
	{"with-hyphen", tokens.NAME, "with-hyphen"},
	{"-starting-with-hyphen", tokens.NAME, "-starting-with-hyphen"},
	{":asdf", tokens.SYMBOL, "asdf"},
	{":as-df", tokens.SYMBOL, "as-df"},
	{"\"\"", tokens.STRING, ""},
	{"\"a\"", tokens.STRING, "a"},
	{"\"a b\"", tokens.STRING, "a b"},
	{"\"abc 12 ! <>?\"", tokens.STRING, "abc 12 ! <>?"},
	{"", tokens.EOF, ""},
}

func TestScanTokens(t *testing.T) {
	for _, tt := range tokenTests {
		c := Scan([]byte(tt.input))
		assertNextToken(t, c, tt.typ, tt.value)
	}
}

func TestScanAfterString(t *testing.T) {
	c := Scan([]byte("\"abc d\" 1"))
	assertNextToken(t, c, tokens.STRING, "abc d")
	assertNextToken(t, c, tokens.NUMBER, "1")
}

func TestScanList(t *testing.T) {
	c := Scan([]byte("[1 2 3]"))

	assertNextToken(t, c, tokens.LBRACKET, "[")
	assertNextToken(t, c, tokens.NUMBER, "1")
	assertNextToken(t, c, tokens.NUMBER, "2")
	assertNextToken(t, c, tokens.NUMBER, "3")
	assertNextToken(t, c, tokens.RBRACKET, "]")
}

func TestScanBlockSimple(t *testing.T) {
	c := Scan([]byte("{1}"))

	assertNextToken(t, c, tokens.LBRACER, "{")
	assertNextToken(t, c, tokens.NUMBER, "1")
	assertNextToken(t, c, tokens.RBRACER, "}")
}

func TestScanBlockMultiple(t *testing.T) {
	c := Scan([]byte("{\n1\n2\n}"))

	assertNextToken(t, c, tokens.LBRACER, "{")
	assertNextToken(t, c, tokens.NUMBER, "1")
	assertNextToken(t, c, tokens.NUMBER, "2")
	assertNextToken(t, c, tokens.RBRACER, "}")
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

func TestScanDot(t *testing.T) {
	c := Scan([]byte("."))

	assertNextToken(t, c, tokens.DOT, ".")
}

func TestScanCommentAtEndOfFile(t *testing.T) {
	c := Scan([]byte("//comment"))

	assertNextToken(t, c, tokens.COMMENT, "comment")
	assertNextToken(t, c, tokens.EOF, "")
}

func TestScanCommentBetweenOtherThings(t *testing.T) {
	c := Scan([]byte("1 // a long comment here\n 2"))

	assertNextToken(t, c, tokens.NUMBER, "1")
	assertNextToken(t, c, tokens.COMMENT, " a long comment here")
	assertNextToken(t, c, tokens.NUMBER, "2")
	assertNextToken(t, c, tokens.EOF, "")
}
