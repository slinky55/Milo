package main

import (
	"fmt"
	"github.com/slinky55/milo/evaluator"
	"github.com/slinky55/milo/lexer"
	"github.com/slinky55/milo/parser"
	"os"
)

func main() {
	b, err := os.ReadFile("./test.milo") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	l := lexer.New(string(b))
	p := parser.New(l)

	program := p.Parse()
	if len(p.Errors) > 0 {
		println("parser had errors")
	}

	e := evaluator.New(program)

	e.Evaluate()
}
