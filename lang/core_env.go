package lang

import (
	"fmt"
)

var nativePrintln *function = MakeFunction(
	MakeArgs("value"),
	MakeNativeBlock(
		func(closure *environment, dynamic *environment) Expr {
			fmt.Println(closure.Lookup("value"))
			return Nil
		}))

func MakeCoreEnv() *environment {
	env := MakeEnv()

	env.Extend("println", nativePrintln)

	return env
}
