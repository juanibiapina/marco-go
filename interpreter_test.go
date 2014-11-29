package marco

import (
	"testing"
)

func TestEvalNumber(t *testing.T) {
	result := Eval(Parse(Scan("1")))
	expected := Number{1}

	if result != expected {
		t.Errorf("Wrong result, expected '%v', got '%v'", 1, result)
	}
}
