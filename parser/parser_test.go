package parser

import (
	"github.com/juanibiapina/marco/lang"
	"github.com/juanibiapina/marco/scanner"
	"github.com/juanibiapina/marco/tokens"
	"reflect"
	"testing"
)

func scan(src string) chan tokens.Token {
	return scanner.Scan([]byte(src))
}

func TestParseNumbers(t *testing.T) {
	block := Parse(scan("1"))

	expected := lang.MakeBlock(lang.MakePair(lang.MakeNumber(1), lang.Nil))

	if !reflect.DeepEqual(block, expected) {
		t.Errorf("Expected '%v' but got '%v'", expected, block)
	}
}

func TestParseNames(t *testing.T) {
	block := Parse(scan("def"))

	expected := lang.MakeBlock(lang.MakePair(lang.MakeName("def"), lang.Nil))

	if !reflect.DeepEqual(block, expected) {
		t.Errorf("Expected '%v' but got '%v'", expected, block)
	}
}

func TestParseString(t *testing.T) {
	block := Parse(scan("\"stuff here\""))

	expected := lang.MakeBlock(
		lang.MakePair(lang.MakeString("stuff here"),
			lang.Nil,
		),
	)

	if !reflect.DeepEqual(block, expected) {
		t.Errorf("Expected '%v' but got '%v'", expected, block)
	}
}

func TestParseTwoNumbers(t *testing.T) {
	block := Parse(scan("1\n\n2"))

	expected := lang.MakeBlock(lang.MakePair(lang.MakeNumber(1), lang.MakePair(lang.MakeNumber(2), lang.Nil)))

	if !reflect.DeepEqual(block, expected) {
		t.Errorf("Expected '%v' but got '%v'", expected, block)
	}
}

func TestParseTwoLists(t *testing.T) {
	block := Parse(scan("[1 2]\n\n[3 4]"))

	expected := lang.MakeBlock(
		lang.MakePair(lang.MakePair(lang.MakeNumber(1),
			lang.MakePair(lang.MakeNumber(2),
				lang.Nil,
			),
		),
			lang.MakePair(lang.MakePair(lang.MakeNumber(3),
				lang.MakePair(lang.MakeNumber(4),
					lang.Nil,
				),
			),
				lang.Nil,
			),
		),
	)

	if !reflect.DeepEqual(block, expected) {
		t.Errorf("Expected '%v' but got '%v'", expected, block)
	}
}

func TestParseApplication(t *testing.T) {
	block := Parse(scan("(a b c)"))

	expected := lang.MakeSingleExprBlock(lang.MakeApplication(lang.SliceToList([]lang.Expr{lang.MakeName("a"), lang.MakeName("b"), lang.MakeName("c")})))

	if !reflect.DeepEqual(block, expected) {
		t.Errorf("Expected '%v' but got '%v'", expected, block)
	}
}

func TestParseNestedNames(t *testing.T) {
	block := Parse(scan("a.b"))

	expected := lang.MakeSingleExprBlock(lang.MakeNestedName("a", lang.MakeName("b")))

	if !reflect.DeepEqual(block, expected) {
		t.Errorf("Expected '%v' but got '%v'", expected, block)
	}
}

func TestParseComments(t *testing.T) {
	block := Parse(scan("a.b // ignored comment"))

	expected := lang.MakeSingleExprBlock(lang.MakeNestedName("a", lang.MakeName("b")))

	if !reflect.DeepEqual(block, expected) {
		t.Errorf("Expected '%v' but got '%v'", expected, block)
	}
}
