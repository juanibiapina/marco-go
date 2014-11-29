package main

import (
	"fmt"
	"github.com/juanibiapina/marco"
)

func main() {
	result := marco.Eval(marco.Parse(marco.Scan("1")))

	fmt.Println(result)
}
