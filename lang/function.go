package lang

type function struct {
	Formal Expr
	Body   Expr
}

func MakeArgs(names ...string) Expr {
	args := make([]Expr, 0, len(names))
	for _, name := range names {
		args = append(args, MakeName(name))
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
			env.Extend(formal.(*name).Value, args[i])
		}
		return body.Invoke(env, dynamic)
	}
	return nil
}

func (f *function) Equal(o Expr) bool {
	switch other := o.(type) {
	case *function:
		return f == other
	default:
		return false
	}
}
