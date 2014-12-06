package lang

type String struct {
	Value string
}

func MakeString(value string) Expr {
	return String{value}
}
