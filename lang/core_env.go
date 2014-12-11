package lang

import (
	"fmt"
)

var nativeEqual *function = MakeFunction(
	MakeArgs("v1", "v2"),
	MakeNativeBlock(
		func(closure *Environment, dynamic *Environment) Expr {
			v1 := closure.Lookup("v1")
			v2 := closure.Lookup("v2")
			return MakeBoolean(v1.Equal(v2))
		}))

var nativeAssert *function = MakeFunction(
	MakeArgs("value"),
	MakeNativeBlock(
		func(closure *Environment, dynamic *Environment) Expr {
			v := closure.Lookup("value")
			if v.(boolean).IsTrue() {
			} else {
				panic("Assertion failed")
			}
			return Nil
		}))

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

	env.Extend("nil", Nil)
	env.Extend("println", nativePrintln)
	env.Extend("def", nativeDef)
	env.Extend("assert", nativeAssert)
	env.Extend("=", nativeEqual)

	return env
}
