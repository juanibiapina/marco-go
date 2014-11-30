package lang

type Number struct {
	Value int64
}

func MakeNumber(v int64) Expr {
	return Number{v}
}
