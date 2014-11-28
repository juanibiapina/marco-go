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

func TestLex(t *testing.T) {
	c := Lex("1")

	select {
	case token := <-c:
		validateToken(t, token, TOKEN_NUMBER, "1")
	case <-time.After(2 * time.Second):
		t.Error("Did not produce a token in time")
	}
}
