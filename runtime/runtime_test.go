package marco

import (
	"github.com/juanibiapina/marco/lang"
	"testing"
)

func TestEvalString(t *testing.T) {
	expr := Eval("1")
	expected := lang.MakeNumber(1)

	if expr != expected {
		t.Errorf("Wrong result, expected '%v', got '%v'", expected, expr)
	}
}

func TestEval(t *testing.T) {
	expr := Eval([]byte("1"))
	expected := lang.MakeNumber(1)

	if expr != expected {
		t.Errorf("Wrong result, expected '%v', got '%v'", expected, expr)
	}
}
