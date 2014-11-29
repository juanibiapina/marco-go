package interpreter

import (
	"github.com/juanibiapina/marco/lang"
)

func Eval(expr lang.Expr, env *lang.Env) lang.Expr {
	switch expr := expr.(type) {
	case lang.Number:
		return expr
	case lang.Name:
		return env.Lookup(expr)
	}
	return nil // error, did not match
}
