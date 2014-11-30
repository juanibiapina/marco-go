package interpreter

import (
	"github.com/juanibiapina/marco/lang"
	"testing"
)

var env *lang.Env

func init() {
	env = lang.MakeEnv()
}

func TestEvalNumber(t *testing.T) {
	result := Eval(lang.MakeNumber(1), env)

	expected := lang.MakeNumber(1)
	if result != expected {
		t.Errorf("Wrong result, expected '%v', got '%v'", expected, result)
	}
}

func TestEvalName(t *testing.T) {
	env.Extend("def", lang.MakeNumber(42))

	result := Eval(lang.Name{"def"}, env)

	expected := lang.MakeNumber(42)
	if result != expected {
		t.Errorf("Wrong result, expected '%v', got '%v'", expected, result)
	}
}

func TestEvalModule(t *testing.T) {
	result := Eval(lang.Module{lang.Pair{lang.MakeNumber(42), lang.MakeNil()}}, env)

	expected := lang.MakeNumber(42)
	if result != expected {
		t.Errorf("Wrong result, expected '%v', got '%v'", expected, result)
	}
}
