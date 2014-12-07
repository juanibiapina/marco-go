package lang

type Nil struct {
}

var nilSingleton Nil

func MakeNil() Expr {
	return nilSingleton
}

func IsNil(e Expr) bool {
	_, ok := e.(Nil)

	return ok
}
