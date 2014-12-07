package lang

func ListToSlice(expr Expr) []Expr {
	var result []Expr
	current := expr
	for !IsNil(current) {
		result = append(result, current.(Pair).First)
		current = current.(Pair).Second
	}
	return result
}

func SliceToList(l []Expr) Expr {
	result := MakeNil()
	for i := len(l) - 1; i >= 0; i-- {
		result = MakePair(l[i], result)
	}
	return result
}
