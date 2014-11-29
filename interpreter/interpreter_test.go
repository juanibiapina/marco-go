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
	result := Eval(lang.Number{1}, env)

	expected := lang.Number{1}
	if result != expected {
		t.Errorf("Wrong result, expected '%v', got '%v'", expected, result)
	}
}

func TestEvalName(t *testing.T) {
	env.Extend("def", lang.Number{42})

	result := Eval(lang.Name{"def"}, env)

	expected := lang.Number{42}
	if result != expected {
		t.Errorf("Wrong result, expected '%v', got '%v'", expected, result)
	}
}
