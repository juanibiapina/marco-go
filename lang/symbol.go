package lang

type Symbol struct {
	Value string
}

func MakeSymbol(name string) Expr {
	return Symbol{name}
}
