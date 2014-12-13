package lang

type function struct {
	Formal Expr
	Body   Expr
}

func MakeArgs(values ...string) Expr {
	args := make([]Expr, 0, len(values))
	for _, v := range values {
		args = append(args, MakeSymbol(v))
	}
	return SliceToList(args)
}

func MakeFunction(formal Expr, body Expr) *function {
	return &function{formal, body}
}

func (f *function) String() string {
	return "Function"
}

func (f *function) Apply(args []Expr, dynamic *Environment) Expr {
	switch body := f.Body.(type) {
	case *nativeBlock:
		env := MakeEnv()
		formals := ListToSlice(f.Formal)
		for i, formal := range formals {
			env.Extend(formal.(*symbol).Value(), args[i])
		}
		return body.Invoke(env, dynamic)
	case *block:
		env := body.Lexical
		formals := ListToSlice(f.Formal)
		for i, formal := range formals {
			env.Extend(formal.(*symbol).Value(), args[i])
		}
		return body.WithEnv(env).Invoke()
	}
	panic("Wrong body type") // TODO better type checker
}

func (f *function) Equal(o Expr) bool {
	switch other := o.(type) {
	case *function:
		return f == other
	default:
		return false
	}
}
