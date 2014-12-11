package lang

import (
	"fmt"
)

type module struct {
	exports map[string]Expr
}

func MakeModule(env *Environment) *module {
	return &module{env.Exports}
}

func (m *module) Lookup(name string) Expr {
	value, ok := m.exports[name]
	if !ok {
		panic(fmt.Sprintf("Module does not export name '%v'", name))
	}
	return value
}

func (m *module) String() string {
	return "Module"
}
