package lang

import (
	"fmt"
)

type number struct {
	value int64
}

func MakeNumber(v int64) *number {
	return &number{v}
}

func (n *number) String() string {
	return fmt.Sprintf("%v", n.value)
}

func (n *number) Equal(o Expr) bool {
	switch other := o.(type) {
	case *number:
		return n.value == other.value
	default:
		return false
	}
}

func (n *number) Plus(o *number) *number {
	return MakeNumber(n.value + o.value)
}

func (n *number) Modulo(o *number) *number {
	return MakeNumber(n.value % o.value)
}
