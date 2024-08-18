package main

import (
	"fmt"
	"github.com/slinky55/milo/lexer"
	"github.com/slinky55/milo/token"
)

func main() {
	input := "let a = 5; var foo = 600;"
	l := lexer.NewLexer(input)

	var t *token.Token
	t = l.NextToken()

	for t.Type != token.EOF {
		fmt.Printf("Type: %s | Literal: %s\n", t.Type, t.Literal)
		t = l.NextToken()
	}
}
