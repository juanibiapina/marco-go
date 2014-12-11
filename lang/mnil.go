package lang

type mnil struct {
}

var Nil mnil = mnil{}

func IsNil(e Expr) bool {
	_, ok := e.(mnil)

	return ok
}

func (n mnil) String() string {
	return "nil"
}

func (n mnil) Equal(o Expr) bool {
	_, ok := o.(mnil)
	return ok
}
