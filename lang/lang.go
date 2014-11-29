package lang

type Expr interface {
}

type Number struct {
	Value int64
}

type Name struct {
	Value string
}

type Env struct {
	bindings map[string]Expr
}

type Error struct {
	Message string
}

func (env *Env) Lookup(name Name) Expr {
	return env.bindings[name.Value]
}

func (env *Env) Extend(name string, expr Expr) {
	env.bindings[name] = expr
}

func MakeEnv() *Env {
	return &Env{
		bindings: make(map[string]Expr),
	}
}
