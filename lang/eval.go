package lang

import (
	"log"
)

func Eval(expr Expr, env *Env) Expr {
	switch expr := expr.(type) {
	case Block:
		return expr.WithEnv(env)
	case Number:
		return expr
	case Symbol:
		return expr
	case String:
		return expr
	case Name:
		return env.Lookup(expr)
	case Nil:
		return expr
	case Pair:
		return MakePair(Eval(expr.First, env), Eval(expr.Second, env))
	default:
		log.Fatalf("Evaluation error, no match for '%v'", expr)
		return nil
	}
}
