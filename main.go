package main

import (
	"os"

  "github.com/slinky55/Milo/token"
)

func main() {
	if len(os.Args) < 2 {
		println("usage: milo [filename]")
		os.Exit(1)
	}

	filename := os.Args[1]

	tokenizer, err := NewTokenizer(filename)
	if err != nil {
		println("Failed to open file: ")
		println(err.Error())
		os.Exit(1)
	}

  for i := 0; i < 10; i++ {
    t := tokenizer.NextToken()
    if t.Type == token.EOF {
      break
    }
    print("Type: " + t.Type)
    println(" | Literal: " + t.Literal)
  }
}
