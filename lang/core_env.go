package lang

import (
	"fmt"
)

var nativeDef *function = MakeFunction(
	MakeArgs("symbol", "value"),
	MakeNativeBlock(
		func(closure *Environment, dynamic *Environment) Expr {
			s := closure.Lookup("symbol")
			v := closure.Lookup("value")
			dynamic.Extend(s.(*symbol).Value(), v)
			return Nil
		}))

var nativePrintln *function = MakeFunction(
	MakeArgs("value"),
	MakeNativeBlock(
		func(closure *Environment, dynamic *Environment) Expr {
			fmt.Println(closure.Lookup("value"))
			return Nil
		}))

func MakeCoreEnv() *Environment {
	env := MakeEnv()

	env.Extend("println", nativePrintln)
	env.Extend("def", nativeDef) // TODO needs tests

	return env
}
