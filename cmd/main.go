package main

import (
	"github.com/slinky55/milo/lexer"
	"github.com/slinky55/milo/parser"
)

func main() {
	input := "let a = 5; var foo = 600;"
	l := lexer.New(input)

	p := parser.New(l)

	p.Parse()
}
