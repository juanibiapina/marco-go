package lang

type Nil struct {
}

func MakeNil() Expr {
	return Nil{}
}
