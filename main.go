package main

import (
	"os"
)

func main() {
	if len(os.Args) < 2 {
		println("usage: milo [filename]")
		os.Exit(1)
	}

	filename := os.Args[1]

	lexer, err := NewLexer(filename)
	if err != nil {
		println("Error creating tokenizer: ")
		println(err.Error())
		os.Exit(1)
	}

  parser := NewParser(lexer) 
  
  ast := parser.GenerateAST()

  print(ast.Name)
}
