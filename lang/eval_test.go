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

	result := Eval(Name{"def"}, env)

	expected := MakeNumber(42)
	if result != expected {
		t.Errorf("Wrong result, expected '%v', got '%v'", expected, result)
	}
}

func TestEvalModule(t *testing.T) {
	result := Eval(Module{Pair{MakeNumber(42), MakeNil()}}, env)

	expected := MakeNumber(42)
	if result != expected {
		t.Errorf("Wrong result, expected '%v', got '%v'", expected, result)
	}
}
