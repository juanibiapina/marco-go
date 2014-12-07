package lang

import (
	"fmt"
)

type symbol struct {
	Value string
}

func MakeSymbol(name string) Expr {
	return &symbol{name}
}

func (s *symbol) String() string {
	return fmt.Sprintf(":%v", s.Value)
}
