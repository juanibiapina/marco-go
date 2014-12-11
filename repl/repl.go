package repl

import (
	"bufio"
	"fmt"
	"github.com/juanibiapina/marco/runtime"
	"os"
)

type Repl struct {
	runtime *runtime.Runtime
}

func New() *Repl {
	return &Repl{runtime.New()}
}

func printGreeting() {
	fmt.Println("Marco Repl")
}

func printPrompt() {
	fmt.Print(":) ")
}

func (r *Repl) Run() {
	scanner := bufio.NewScanner(os.Stdin)

	printGreeting()

	for {
		printPrompt()
		scanner.Scan()
		line := scanner.Text() // TODO check for scanner errors
		result := r.runtime.Run(line)
		fmt.Println(result)
	}
}
