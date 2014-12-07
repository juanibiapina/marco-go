package lang

type Block struct {
	Forms   Expr
	Lexical *Env
}

func MakeBlock(forms []Expr) Block {
	return Block{SliceToList(forms), nil}
}

func MakeSingleExprBlock(form Expr) Block {
	return MakeBlock([]Expr{form})
}

func ToBlock(expr Expr) Block {
	block, ok := expr.(Block)

	if !ok {
		panic("Expected Block") // TODO better error handling
	}

	return block
}

func (block Block) WithEnv(env *Env) Block {
	return Block{block.Forms, env}
}

func (block Block) Invoke() Expr {
	forms := ListToSlice(block.Forms)

	var result Expr = MakeNil()
	for _, form := range forms {
		result = Eval(form, block.Lexical)
	}
	return result
}
