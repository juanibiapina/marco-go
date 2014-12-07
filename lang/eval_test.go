package lang

import (
	"testing"
)

var env *Env

func init() {
	env = MakeEnv()
}

func TestEvalNumber(t *testing.T) {
	result := Eval(MakeNumber(1), env)

	expected := MakeNumber(1)
	if result != expected {
		t.Errorf("Wrong result, expected '%v', got '%v'", expected, result)
	}
}

func TestEvalName(t *testing.T) {
	env.Extend("def", MakeNumber(42))

	result := Eval(MakeName("def"), env)

	expected := MakeNumber(42)
	if result != expected {
		t.Errorf("Wrong result, expected '%v', got '%v'", expected, result)
	}
}

func TestEvalBlock(t *testing.T) {
	result := Eval(MakeSingleExprBlock(MakeNumber(42)), env)

	expected := Block{MakePair(MakeNumber(42), MakeNil()), env}

	if result != expected {
		t.Errorf("Wrong result, expected '%v', got '%v'", expected, result)
	}
}

func TestEvalString(t *testing.T) {
	result := Eval(MakeString("some string"), env)

	expected := MakeString("some string")
	if result != expected {
		t.Errorf("Wrong result, expected '%v', got '%v'", expected, result)
	}
}
