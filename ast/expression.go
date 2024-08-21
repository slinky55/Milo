package ast

import "github.com/slinky55/milo/token"

type Expression interface {
	Node
	expressionNode()
}

type IdentExpr struct {
	Token *token.Token
	Value string
}

func (ie *IdentExpr) Literal() string {
	return ie.Token.Literal
}

func (ie *IdentExpr) expressionNode() { /* EMPTY */ }

type NumberExpr struct {
	Token *token.Token
	Value float64
}

func (ne *NumberExpr) Literal() string {
	return ne.Token.Literal
}

func (ne *NumberExpr) expressionNode() { /* EMPTY */ }
