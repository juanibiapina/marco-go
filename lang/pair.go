package lang

import (
	"fmt"
)

type pair struct {
	First  Expr
	Second Expr
}

func MakePair(e1 Expr, e2 Expr) *pair {
	return &pair{e1, e2}
}

func (p *pair) String() string {
	return fmt.Sprintf("[%v %v]", p.First.String(), p.Second.String())
}

func (p *pair) Equal(o Expr) bool {
	switch other := o.(type) {
	case *pair:
		return p.First == other.First && p.Second == other.Second
	default:
		return false
	}
}
