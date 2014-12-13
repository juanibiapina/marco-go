package lang

import (
	"fmt"
)

type Environment struct {
	bindings map[string]Expr
	Exports  map[string]Expr
	Parent   *Environment
}

func MakeEnv() *Environment {
	return &Environment{
		bindings: make(map[string]Expr),
		Exports:  make(map[string]Expr),
	}
}

func SpawnEnv(parent *Environment) *Environment {
	env := MakeEnv()
	env.Parent = parent
	return env
}

func (env *Environment) Lookup(name string) Expr {
	value, ok := env.bindings[name]
	if !ok {
		if env.Parent != nil {
			return env.Parent.Lookup(name)
		} else {
			panic(fmt.Sprintf("Binding not found: '%v'", name))
		}
	}
	return value
}

func (env *Environment) Extend(name string, expr Expr) {
	env.bindings[name] = expr
}

func (env *Environment) Export(name string) {
	value := env.Lookup(name)
	env.Exports[name] = value
}

func (env *Environment) String() string {
	return "Environment"
}

func (env *Environment) Equal(o Expr) bool {
	return false
}
