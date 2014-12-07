package marco

import (
	"github.com/juanibiapina/marco/lang"
	"reflect"
	"testing"
)

func TestRunString(t *testing.T) {
	expr := Run("1")
	expected := lang.MakeNumber(1)

	if !reflect.DeepEqual(expr, expected) {
		t.Errorf("Wrong result, expected '%v', got '%v'", expected, expr)
	}
}
