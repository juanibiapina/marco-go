package lang

import (
	"fmt"
)

type pair struct {
	First  Expr
	Second Expr
}

func MakePair(e1 Expr, e2 Expr) Expr {
	return &pair{e1, e2}
}

func (p *pair) String() string {
	return fmt.Sprintf("[%v %v]", p.First.String(), p.Second.String())
}
