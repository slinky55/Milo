package main

type Parser struct {
  lexer *Lexer 
}

type AST interface



func NewParser(l *Lexer) *Parser {
  return &Parser{
    lexer: l,
  }
}

func (l *Parser) GenerateAST() *AST {
  return &AST {
    Name: "main",
  }
}
