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
	block := Parse(scan("1"))

	expected := lang.Block{lang.Pair{lang.MakeNumber(1), lang.MakeNil()}, nil}

	if block != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, block)
	}
}

func TestParseNames(t *testing.T) {
	block := Parse(scan("def"))

	expected := lang.Block{lang.Pair{lang.Name{"def"}, lang.MakeNil()}, nil}

	if block != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, block)
	}
}

func TestParseString(t *testing.T) {
	block := Parse(scan("\"stuff here\""))

	expected := lang.Block{
		lang.Pair{
			lang.String{"stuff here"},
			lang.MakeNil(),
		},
		nil,
	}

	if block != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, block)
	}
}

func TestParseTwoNumbers(t *testing.T) {
	block := Parse(scan("1\n\n2"))

	expected := lang.Block{lang.Pair{lang.MakeNumber(1), lang.Pair{lang.MakeNumber(2), lang.MakeNil()}}, nil}

	if block != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, block)
	}
}

func TestParseTwoLists(t *testing.T) {
	block := Parse(scan("[1 2]\n\n[3 4]"))

	expected := lang.Block{
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
		nil,
	}

	if block != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, block)
	}
}

func TestParseApplication(t *testing.T) {
	block := Parse(scan("(a b c)"))

	expected := lang.MakeSingleExprBlock(lang.Application{
		lang.SliceToList([]lang.Expr{lang.MakeName("a"), lang.MakeName("b"), lang.MakeName("c")}),
	})

	if block != expected {
		t.Errorf("Expected '%v' but got '%v'", expected, block)
	}
}
