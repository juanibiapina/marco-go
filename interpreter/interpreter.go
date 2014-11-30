package interpreter

import (
	"github.com/juanibiapina/marco/lang"
	"log"
)

func Eval(expr lang.Expr, env *lang.Env) lang.Expr {
	switch expr := expr.(type) {
	case lang.Number:
		return expr
	case lang.Name:
		return env.Lookup(expr)
	case lang.Nil:
		return expr
	case lang.Pair:
		return lang.MakePair(Eval(expr.First, env), Eval(expr.Second, env))
	default:
		log.Fatalf("Evaluation error, no match for '%v'", expr)
		return nil
	}
}
