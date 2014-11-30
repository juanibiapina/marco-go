package lang

func MakeList(l []Expr) Expr {
	result := MakeNil()
	for i := len(l) - 1; i >= 0; i-- {
		result = MakePair(l[i], result)
	}
	return result
}
