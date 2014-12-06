package lang

type Name struct {
	Value string
}

func MakeName(value string) Expr {
	return Name{value}
}
