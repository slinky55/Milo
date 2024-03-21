package main

import (
	"os"

  "github.com/slinky55/milo/lexer"
  "github.com/slinky55/milo/parser"
)

func main() {
	if len(os.Args) < 2 {
		println("usage: milo [filename]")
		os.Exit(1)
	}

	filename := os.Args[1]

	lexer, err := lexer.New(filename)
	if err != nil {
		println("Error creating tokenizer: ")
		println(err.Error())
		os.Exit(1)
	}

  parser := parser.New(lexer) 
  
  _ = parser.ParseProgram()
}
