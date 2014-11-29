package parser

import (
	"github.com/juanibiapina/marco/lang"
	"github.com/juanibiapina/marco/scanner"
	"github.com/juanibiapina/marco/tokens"
	"testing"
)

func scan(src string) chan tokens.Token {
	return scanner.Scan([]byte(src))
}

func TestParseNumbers(t *testing.T) {
	ast := Parse(scan("1"))

	v, ok := ast.(lang.Number)

	if !ok {
		t.Error("Expected type Number")
	}
	if v.Value != 1 {
		t.Errorf("Expected '%v' but got '%v'", 1, v.Value)
	}
}
