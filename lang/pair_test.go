package lang

import (
	"testing"
)

func TestPairString(t *testing.T) {
	result := MakePair(MakeString("a"), MakeNumber(2)).String()

	expected := "[\"a\" 2]"

	if result != expected {
		t.Errorf("Wrong result, expected '%v', got '%v'", expected, result)
	}
}
