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
