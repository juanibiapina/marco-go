package marco

import (
	"github.com/juanibiapina/marco/interpreter"
	"github.com/juanibiapina/marco/lang"
	"github.com/juanibiapina/marco/parser"
	"github.com/juanibiapina/marco/scanner"
)

func makeSourceString(src string) []byte {
	return []byte(src)
}

func Eval(src []byte) lang.Expr {
	tokens := scanner.Scan(src)
	ast := parser.Parse(tokens)
	expr := interpreter.Eval(ast)
	return expr
}

func EvalString(src string) lang.Expr {
	tokens := scanner.Scan(makeSourceString(src))
	ast := parser.Parse(tokens)
	expr := interpreter.Eval(ast)
	return expr
}
