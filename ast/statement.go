package ast

import "github.com/slinky55/milo/token"

type Statement interface {
	Node
	statementNode()
}

type LetStatement struct {
	Token *token.Token
	Ident *IdentExpr
	Expr  Expression
}

func (ls *LetStatement) Literal() string {
	return ls.Token.Literal
}

func (ls *LetStatement) statementNode() { /* EMPTY */ }

type ReturnStatement struct {
	Token *token.Token
	Expr  Expression
}

func (rs *ReturnStatement) Literal() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) statementNode() { /* EMPTY */ }
