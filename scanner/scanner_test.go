package scanner

import (
	"github.com/juanibiapina/marco/tokens"
	"testing"
	"time"
)

func validateToken(t *testing.T, token tokens.Token, typ tokens.TokenType, value string) {
	if token.Typ != typ {
		t.Errorf("Wrong token type: expected '%v', got '%v'", typ, token.Typ)
	}
	if token.Value != value {
		t.Errorf("Wrong token value: expected '%v', got '%v'", value, token.Value)
	}
}

func assertNextToken(t *testing.T, c chan tokens.Token, typ tokens.TokenType, value string) {
	select {
	case token := <-c:
		validateToken(t, token, tokens.NUMBER, value)
	case <-time.After(2 * time.Second):
		t.Errorf("Expected '%v' but did not produce a token in time", value)
	}
}

var numbersTests = []struct {
	value string
	typ   tokens.TokenType
}{
	{"1", tokens.NUMBER},
	{"2", tokens.NUMBER},
	{"823", tokens.NUMBER},
}

func TestScanNumbers(t *testing.T) {
	for _, tt := range numbersTests {
		c := Scan([]byte(tt.value))
		assertNextToken(t, c, tt.typ, tt.value)
	}
}
