package lang

type Expr interface {
}

type Number struct {
	Value int64
}

type Error struct {
	Message string
}
