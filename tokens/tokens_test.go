package tokens

import (
	"fmt"
	"testing"
)

func TestTokenTypeToString(t *testing.T) {
	typ := NAME

	result := fmt.Sprint(typ)
	expected := "TOKEN_NAME"
	if result != expected {
		t.Errorf("Wrong result, expected '%v', got '%v'", expected, result)
	}
}
