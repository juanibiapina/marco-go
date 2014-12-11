package lang

type block struct {
	Forms   Expr
	Lexical *Environment
}

func MakeBlock(forms Expr) *block {
	return &block{forms, nil}
}

func MakeSingleExprBlock(form Expr) *block {
	return MakeBlock(SliceToList([]Expr{form}))
}

func ToBlock(expr Expr) *block {
	block, ok := expr.(*block)

	if !ok {
		panic("Expected block") // TODO better error handling
	}

	return block
}

func (b *block) WithEnv(env *Environment) *block {
	return &block{b.Forms, env}
}

func (block *block) Invoke() Expr {
	forms := ListToSlice(block.Forms)

	var result Expr = Nil
	for _, form := range forms {
		result = Eval(form, block.Lexical)
	}
	return result
}

func (block *block) String() string {
	return "Block"
}

func (b *block) Equal(o Expr) bool {
	switch other := o.(type) {
	case *block:
		return b == other
	default:
		return false
	}
}
