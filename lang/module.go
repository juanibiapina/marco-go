package lang

type Module struct {
	Forms Expr
}

func MakeModule(forms []Expr) Expr {
	return Module{MakeList(forms)}
}
