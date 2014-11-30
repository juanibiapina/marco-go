package lang

type Pair struct {
	First  Expr
	Second Expr
}

func MakePair(e1 Expr, e2 Expr) Expr {
	return Pair{e1, e2}
}
