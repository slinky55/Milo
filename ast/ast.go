package ast

type AST interface {
	TokenLit() string
}

type Program struct {
	Statements []Statement
}
