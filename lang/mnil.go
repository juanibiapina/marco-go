package lang

type mnil struct {
}

var nilSingleton mnil

func MakeNil() Expr {
	return &nilSingleton
}

func IsNil(e Expr) bool {
	_, ok := e.(*mnil)

	return ok
}

func (n *mnil) String() string {
	return "nil"
}
