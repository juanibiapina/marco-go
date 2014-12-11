package lang

import (
	"fmt"
)

type Expr interface {
	fmt.Stringer
	Equal(Expr) bool
}
