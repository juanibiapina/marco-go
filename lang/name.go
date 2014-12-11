package lang

import (
	"fmt"
)

type name struct {
	Value  string
	Nested Expr
}

func MakeName(value string) Expr {
	return &name{value, nil}
}

func MakeNestedName(value string, nested Expr) Expr {
	return &name{value, nested}
}

func (n *name) String() string {
	if n.Nested != nil {
		return fmt.Sprintf("%v.%v", n.Value, n.Nested.String())
	}
	return n.Value
}

func (s *name) Equal(o Expr) bool {
	return false
}
