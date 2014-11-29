package marco

import (
	"github.com/juanibiapina/marco/lang"
	"testing"
)

func TestEvalString(t *testing.T) {
	expr := EvalString("1")
	expected := lang.Number{1}

	if expr != expected {
		t.Errorf("Wrong result, expected '%v', got '%v'", expected, expr)
	}
}
