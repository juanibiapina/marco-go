package lang

import (
	"fmt"
)

type number struct {
	Value int64
}

func MakeNumber(v int64) *number {
	return &number{v}
}

func (n *number) String() string {
	return fmt.Sprintf("%v", n.Value)
}
