package marco

import (
	"testing"
)

func TestParseNumbers(t *testing.T) {
	ast := Parse(Scan("1"))

	v, ok := ast.(Number)

	if !ok {
		t.Error("Expected type Number")
	}
	if v.value != 1 {
		t.Errorf("Expected '%v' but got '%v'", 1, v.value)
	}
}
