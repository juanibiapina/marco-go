package lang

type Expr interface {
}

type Number struct {
	Value int64
}

type Name struct {
	Value string
}

type Error struct {
	Message string
}
