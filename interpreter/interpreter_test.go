package interpreter

import (
	"github.com/juanibiapina/marco/lang"
	"testing"
)

func TestEvalNumber(t *testing.T) {
	result := Eval(lang.Number{1})
	expected := lang.Number{1}

	if result != expected {
		t.Errorf("Wrong result, expected '%v', got '%v'", 1, result)
	}
}
