package lang

import (
	"reflect"
	"testing"
)

func TestListToSlice(t *testing.T) {
	list := MakePair(MakeNumber(42), MakePair(MakeString("value"), Nil))

	result := ListToSlice(list)
	expected := []Expr{MakeNumber(42), MakeString("value")}

	if len(result) != len(expected) {
		t.Errorf("Result has wrong length, expected '%v', got '%v'", len(expected), len(result))
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wrong result, expected '%v', got '%v'", expected, result)
	}
}

func TestSliceToList(t *testing.T) {
	list := []Expr{MakeNumber(42), MakeString("value")}

	result := SliceToList(list)
	expected := MakePair(MakeNumber(42), MakePair(MakeString("value"), Nil))

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wrong result, expected '%v', got '%v'", expected, result)
	}
}
