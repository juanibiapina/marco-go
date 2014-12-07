package lang

import (
	"fmt"
)

type mstring struct {
	value string
}

func MakeString(value string) Expr {
	return &mstring{value}
}

func (s *mstring) String() string {
	return fmt.Sprintf("\"%v\"", s.value)
}
