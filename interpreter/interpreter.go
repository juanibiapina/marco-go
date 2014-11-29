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
	default:
		log.Fatalf("Evaluation error, no match for '%v'", expr)
		return nil
	}
}
