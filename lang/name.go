package lang

type Name struct {
	Value  string
	Nested Expr
}

func MakeName(value string) Expr {
	return Name{value, nil}
}

func MakeNestedName(value string, nested Expr) Expr {
	return Name{value, nested}
}
