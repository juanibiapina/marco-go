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

	module, ok := ast.(lang.Module)

	if !ok {
		t.Error("Wrong type, expected 'lang.Module', got '%T'", ast)
	}

	expected := lang.Module{lang.Pair{lang.MakeNumber(1), lang.MakeNil()}}

	if module != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, module)
	}
}

func TestParseNames(t *testing.T) {
	ast := Parse(scan("def"))

	module, ok := ast.(lang.Module)

	if !ok {
		t.Error("Wrong type, expected 'lang.Module', got '%T'", ast)
	}

	expected := lang.Module{lang.Pair{lang.Name{"def"}, lang.MakeNil()}}

	if module != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, module)
	}
}

func TestParseTwoNumbers(t *testing.T) {
	ast := Parse(scan("1\n\n2"))

	module, ok := ast.(lang.Module)

	if !ok {
		t.Error("Wrong type, expected 'lang.Module', got '%T'", ast)
	}

	expected := lang.Module{lang.Pair{lang.MakeNumber(1), lang.Pair{lang.MakeNumber(2), lang.MakeNil()}}}

	if module != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, module)
	}
}

func TestParseTwoLists(t *testing.T) {
	ast := Parse(scan("[1 2]\n\n[3 4]"))

	module, ok := ast.(lang.Module)

	if !ok {
		t.Error("Wrong type, expected 'lang.Module', got '%T'", ast)
	}

	expected := lang.Module{
		lang.Pair{
			lang.Pair{
				lang.MakeNumber(1),
				lang.Pair{
					lang.MakeNumber(2),
					lang.MakeNil(),
				},
			},
			lang.Pair{
				lang.Pair{
					lang.MakeNumber(3),
					lang.Pair{
						lang.MakeNumber(4),
						lang.MakeNil(),
					},
				},
				lang.MakeNil(),
			},
		},
	}

	if module != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, module)
	}
}
