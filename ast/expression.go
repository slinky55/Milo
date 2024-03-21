package ast

type Expression interface {
	AST
	expression()
}
