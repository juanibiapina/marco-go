package marco

import (
	"github.com/juanibiapina/marco/lang"
)

func Run(src interface{}) lang.Expr {
	return run(convertInput(src))
}
