package lang

import (
	"fmt"
)

type environment struct {
	bindings map[string]Expr
	Exports  map[string]Expr
}

func MakeEnv() *environment {
	return &environment{
		bindings: make(map[string]Expr),
		Exports:  make(map[string]Expr),
	}
}

func (env *environment) Lookup(name string) Expr {
	value, ok := env.bindings[name]
	if !ok {
		panic(fmt.Sprintf("Binding not found: '%v'", name))
	}
	return value
}

func (env *environment) Extend(name string, expr Expr) {
	env.bindings[name] = expr
}

func (env *environment) Export(name string) {
	value := env.Lookup(name)
	env.Exports[name] = value
}

func (env *environment) String() string {
	return "Environment"
}
