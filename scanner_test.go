package marco

import (
	"testing"
	"time"
)

func validateToken(t *testing.T, token Token, typ TokenType, value string) {
	if token.typ != typ {
		t.Errorf("Wrong token type: expected '%v', got '%v'", typ, token.typ)
	}
	if token.value != value {
		t.Errorf("Wrong token value: expected '%v', got '%v'", value, token.value)
	}
}

func assertNextToken(t *testing.T, c chan Token, typ TokenType, value string) {
	select {
	case token := <-c:
		validateToken(t, token, TOKEN_NUMBER, value)
	case <-time.After(2 * time.Second):
		t.Errorf("Expected '%v' but did not produce a token in time", value)
	}
}

var numbersTests = []struct {
	value string
	typ   TokenType
}{
	{"1", TOKEN_NUMBER},
	{"2", TOKEN_NUMBER},
	{"823", TOKEN_NUMBER},
}

func TestScanNumbers(t *testing.T) {
	for _, tt := range numbersTests {
		c := Scan(tt.value)
		assertNextToken(t, c, tt.typ, tt.value)
	}
}
