package lang

func ListToSlice(expr Expr) []Expr {
	var result []Expr
	current := expr
	for !IsNil(current) {
		result = append(result, current.(*pair).First)
		current = current.(*pair).Second
	}
	return result
}

func SliceToList(l []Expr) Expr {
	var result Expr = Nil
	for i := len(l) - 1; i >= 0; i-- {
		result = MakePair(l[i], result)
	}
	return result
}
