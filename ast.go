package marco

type Expr interface {
}

type Number struct {
	value int64
}

type Error struct {
	message string
}
