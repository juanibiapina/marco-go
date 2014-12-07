package lang

type environment struct {
	bindings map[string]Expr
}

func MakeEnv() *environment {
	return &environment{
		bindings: make(map[string]Expr),
	}
}

func (env *environment) Lookup(name *name) Expr {
	return env.bindings[name.Value]
}

func (env *environment) Extend(name string, expr Expr) {
	env.bindings[name] = expr
}

func (env *environment) String() string {
	return "TODO"
}
