package lang

import (
	"fmt"
)

type symbol struct {
	value string
}

func MakeSymbol(name string) Expr {
	return &symbol{name}
}

func (s *symbol) String() string {
	return fmt.Sprintf(":%v", s.value)
}

func (s *symbol) Value() string {
	return s.value
}

func (s *symbol) Equal(o Expr) bool {
	switch other := o.(type) {
	case *symbol:
		return s.value == other.value
	default:
		return false
	}
}
