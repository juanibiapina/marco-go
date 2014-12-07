package marco

import (
	"github.com/juanibiapina/marco/lang"
	"testing"
)

func TestRunString(t *testing.T) {
	expr := Run("1")
	expected := lang.MakeNumber(1)

	if expr != expected {
		t.Errorf("Wrong result, expected '%v', got '%v'", expected, expr)
	}
}
