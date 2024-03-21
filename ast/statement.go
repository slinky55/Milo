package ast

import "github.com/slinky55/milo/token"

type Statement interface {
	AST
	statement()
}

type LetStatement struct {
	Token *token.Token
	Name  Expression
	Value Expression
}

func (ls *LetStatement) TokenLit() string {
	return "let"
}

func (ls *LetStatement) statement() {}

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (rs *ReturnStatement) TokenLit() string {
	return "return"
}

func (rs *ReturnStatement) statement() {}
