package lang

import (
	"fmt"
)

type environment struct {
	bindings map[string]Expr
}

func MakeEnv() *environment {
	return &environment{
		bindings: make(map[string]Expr),
	}
}

func (env *environment) Lookup(name *name) Expr {
	value, ok := env.bindings[name.Value]
	if !ok {
		panic(fmt.Sprintf("Binding not found: '%v'", name))
	}
	return value
}

func (env *environment) Extend(name string, expr Expr) {
	env.bindings[name] = expr
}

func (env *environment) String() string {
	return "Environment"
}
