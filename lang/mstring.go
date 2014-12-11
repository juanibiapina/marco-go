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

func (s *mstring) Equal(o Expr) bool {
	switch other := o.(type) {
	case *mstring:
		return s.value == other.value
	default:
		return false
	}
}
