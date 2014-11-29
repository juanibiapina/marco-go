package marco

import (
	"github.com/juanibiapina/marco/interpreter"
	"github.com/juanibiapina/marco/lang"
	"github.com/juanibiapina/marco/parser"
	"github.com/juanibiapina/marco/scanner"
)

func makeSource(src string) []byte {
	return []byte(src)
}

func EvalString(src string) lang.Expr {
	tokens := scanner.Scan(makeSource(src))
	ast := parser.Parse(tokens)
	expr := interpreter.Eval(ast)
	return expr
}
