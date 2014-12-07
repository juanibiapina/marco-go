package marco

import (
	"github.com/juanibiapina/marco/lang"
	"github.com/juanibiapina/marco/parser"
	"github.com/juanibiapina/marco/scanner"
	"log"
)

func convertInput(src interface{}) []byte {
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

func run(src []byte) lang.Expr {
	tokens := scanner.Scan(src)
	blockAst := parser.Parse(tokens)
	env := lang.MakeEnv()
	expr := lang.Eval(blockAst, env)
	block := lang.ToBlock(expr)
	return block.Invoke()
}

func Run(src interface{}) lang.Expr {
	return run(convertInput(src))
}
