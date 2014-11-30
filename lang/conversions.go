package lang

func ModuleToSlice(expr Expr) []Expr {
	var result []Expr
	current := expr.(Module).Forms
	for !IsNil(current) {
		result = append(result, current.(Pair).First)
		current = current.(Pair).Second
	}
	return result
}
