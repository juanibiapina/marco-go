package lang

type Env struct {
	bindings map[string]Expr
}

func MakeEnv() *Env {
	return &Env{
		bindings: make(map[string]Expr),
	}
}

func (env *Env) Lookup(name Name) Expr {
	return env.bindings[name.Value]
}

func (env *Env) Extend(name string, expr Expr) {
	env.bindings[name] = expr
}
