package main

import (
	"flag"
	"fmt"
	"github.com/juanibiapina/marco/runtime"
	"io/ioutil"
	"log"
	"os"
)

func usage() {
	fmt.Println("Usage: marco <filename>")
	os.Exit(1)
}

func getFileName() string {
	if len(flag.Args()) != 1 {
		usage()
	}
	return flag.Args()[0]
}

func parseArguments() {
	flag.Parse()
}

func main() {
	parseArguments()
	filename := getFileName()

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file '%v': %v", filename, err)
	}

	r := runtime.New()
	result := r.Run(data)

	fmt.Println(result)
}
