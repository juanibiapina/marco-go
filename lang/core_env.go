package lang

import (
	"fmt"
)

var nativeCons *function = MakeFunction(
	MakeArgs("head", "tail"),
	MakeNativeBlock(
		func(closure *Environment, dynamic *Environment) Expr {
			head := closure.Lookup("head")
			tail := closure.Lookup("tail")

			return MakePair(head, tail)
		}))

var nativeIf *function = MakeFunction(
	MakeArgs("cond", "thenClause", "elseClause"),
	MakeNativeBlock(
		func(closure *Environment, dynamic *Environment) Expr {
			cond := closure.Lookup("cond")
			thenClause := closure.Lookup("thenClause")
			elseClause := closure.Lookup("elseClause")

			if cond.(boolean).IsTrue() {
				return thenClause.(*block).Invoke()
			} else {
				return elseClause.(*block).Invoke()
			}
		}))

var nativeFunction *function = MakeFunction(
	MakeArgs("formal", "body"),
	MakeNativeBlock(
		func(closure *Environment, dynamic *Environment) Expr {
			formal := closure.Lookup("formal")
			body := closure.Lookup("body")
			return MakeFunction(formal, body)
		}))

var nativeEqual *function = MakeFunction(
	MakeArgs("v1", "v2"),
	MakeNativeBlock(
		func(closure *Environment, dynamic *Environment) Expr {
			v1 := closure.Lookup("v1")
			v2 := closure.Lookup("v2")
			return MakeBoolean(v1.Equal(v2))
		}))

var nativePlus *function = MakeFunction(
	MakeArgs("v1", "v2"),
	MakeNativeBlock(
		func(closure *Environment, dynamic *Environment) Expr {
			v1 := closure.Lookup("v1")
			v2 := closure.Lookup("v2")
			return v1.(*number).Plus(v2.(*number))
		}))

var nativeModulo *function = MakeFunction(
	MakeArgs("v1", "v2"),
	MakeNativeBlock(
		func(closure *Environment, dynamic *Environment) Expr {
			v1 := closure.Lookup("v1")
			v2 := closure.Lookup("v2")
			return v1.(*number).Modulo(v2.(*number))
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

	env.Extend("assert", nativeAssert)

	env.Extend("nil", Nil)

	env.Extend("true", MakeBoolean(true))
	env.Extend("false", MakeBoolean(false))

	env.Extend("if", nativeIf)

	env.Extend("println", nativePrintln)

	env.Extend("def", nativeDef)

	env.Extend("=", nativeEqual)

	env.Extend("+", nativePlus)
	env.Extend("%", nativeModulo)

	env.Extend("function", nativeFunction)

	env.Extend("cons", nativeCons)

	return env
}
