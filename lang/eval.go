package lang

import (
	"log"
)

func Eval(expr Expr, env *environment) Expr {
	switch expr := expr.(type) {
	case *block:
		return expr.WithEnv(env)
	case *number:
		return expr
	case *symbol:
		return expr
	case *mstring:
		return expr
	case *name:
		result := env.Lookup(expr.Value)
		nextLevel := expr.Nested
		if nextLevel != nil {
			result = result.(*module).Lookup(nextLevel.(*name).Value)
			nextLevel = nextLevel.(*name).Nested
		}
		return result
	case *mnil:
		return expr
	case *pair:
		return MakePair(Eval(expr.First, env), Eval(expr.Second, env))
	default:
		log.Fatalf("Evaluation error, no match for '%v'", expr)
		return nil
	}
}
