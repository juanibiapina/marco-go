package marco

import (
	"github.com/juanibiapina/marco/interpreter"
	"github.com/juanibiapina/marco/lang"
	"github.com/juanibiapina/marco/parser"
	"github.com/juanibiapina/marco/scanner"
	"log"
)

func convert(src interface{}) []byte {
	switch src := src.(type) {
	case []byte:
		return src
	case string:
		return []byte(src)
	default:
		log.Fatalf("Unnexpected type: '%T'", src) // use errors instead
		return nil
	}
}

func eval(src []byte) lang.Expr {
	tokens := scanner.Scan(src)
	ast := parser.Parse(tokens)
	expr := interpreter.Eval(ast, lang.MakeEnv())
	return expr
}

func Eval(src interface{}) lang.Expr {
	return eval(convert(src))
}
