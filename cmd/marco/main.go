package main

import (
	"fmt"
	"github.com/juanibiapina/marco/interpreter"
	"github.com/juanibiapina/marco/parser"
	"github.com/juanibiapina/marco/scanner"
)

func main() {
	result := interpreter.Eval(parser.Parse(scanner.Scan("1")))

	fmt.Println(result)
}
