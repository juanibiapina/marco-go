package lang

type Nil struct {
}

func MakeNil() Expr {
	return Nil{}
}

func IsNil(e Expr) bool {
	_, ok := e.(Nil)

	return ok
}
