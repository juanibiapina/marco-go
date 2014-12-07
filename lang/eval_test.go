package lang

import (
	"reflect"
	"testing"
)

var env *environment

func init() {
	env = MakeEnv()
}

func TestEvalNumber(t *testing.T) {
	result := Eval(MakeNumber(1), env)

	expected := MakeNumber(1)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wrong result, expected '%v', got '%v'", expected, result)
	}
}

func TestEvalName(t *testing.T) {
	env.Extend("def", MakeNumber(42))

	result := Eval(MakeName("def"), env)

	expected := MakeNumber(42)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wrong result, expected '%v', got '%v'", expected, result)
	}
}

func TestEvalNestedName(t *testing.T) {
	moduleEnv := MakeEnv()
	moduleEnv.Extend("b", MakeNumber(1))
	moduleEnv.Export("b")

	env.Extend("a", MakeModule(moduleEnv))

	result := Eval(MakeNestedName("a", MakeName("b")), env)

	expected := MakeNumber(1)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wrong result, expected '%v', got '%v'", expected, result)
	}
}

func TestEvalBlock(t *testing.T) {
	result := Eval(MakeSingleExprBlock(MakeNumber(42)), env)

	expected := MakeBlock(MakePair(MakeNumber(42), Nil)).WithEnv(env)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wrong result, expected '%v', got '%v'", expected, result)
	}
}

func TestEvalString(t *testing.T) {
	result := Eval(MakeString("some string"), env)

	expected := MakeString("some string")
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Wrong result, expected '%v', got '%v'", expected, result)
	}
}
