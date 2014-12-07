package lang

import (
	"testing"
)

func TestNestedNameString(t *testing.T) {
	result := MakeNestedName("a", MakeName("b")).String()

	expected := "a.b"

	if result != expected {
		t.Errorf("Wrong result, expected '%v', got '%v'", expected, result)
	}
}
