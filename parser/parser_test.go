package parser

import (
	"github.com/juanibiapina/marco/lang"
	"github.com/juanibiapina/marco/scanner"
	"testing"
)

func TestParseNumbers(t *testing.T) {
	ast := Parse(scanner.Scan([]byte("1")))

	v, ok := ast.(lang.Number)

	if !ok {
		t.Error("Expected type Number")
	}
	if v.Value != 1 {
		t.Errorf("Expected '%v' but got '%v'", 1, v.Value)
	}
}
