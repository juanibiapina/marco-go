package runtime

import (
	"github.com/juanibiapina/marco/lang"
	"github.com/juanibiapina/marco/parser"
	"github.com/juanibiapina/marco/scanner"
	"log"
)

type Runtime struct {
	env *lang.Environment
}

func New() *Runtime {
	return &Runtime{lang.MakeCoreEnv()}
}

func convertInput(src interface{}) []byte {
	switch src := src.(type) {
	case []byte:
		return src
	case string:
		return []byte(src)
	default:
		log.Fatalf("Unnexpected type: '%T'", src) // TODO use errors instead
		return nil
	}
}

func (r *Runtime) Run(isrc interface{}) lang.Expr {
	src := convertInput(isrc)
	tokens := scanner.Scan(src)
	blockAst := parser.Parse(tokens)
	expr := lang.Eval(blockAst, r.env)
	block := lang.ToBlock(expr)
	return block.Invoke()
}
