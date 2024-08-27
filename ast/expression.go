package ast

import (
	"github.com/slinky55/milo/token"
	"strings"
)

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

func (ie *IdentExpr) ToString() string {
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

func (ne *NumberExpr) ToString() string {
	return ne.Token.Literal
}

func (ne *NumberExpr) expressionNode() { /* EMPTY */ }

type PrefixExpression struct {
	Token    *token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) Literal() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) ToString() string {
	var out strings.Builder

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.ToString())
	out.WriteString(")")

	return out.String()
}

func (pe *PrefixExpression) expressionNode() { /* EMPTY */ }

type BinaryExpression struct {
	Token    *token.Token
	Operator string
	Left     Expression
	Right    Expression
}

func (be *BinaryExpression) Literal() string {
	return be.Token.Literal
}

func (be *BinaryExpression) ToString() string {
	var out strings.Builder

	out.WriteString("(")
	out.WriteString(be.Left.ToString())
	out.WriteString(" " + be.Operator + " ")
	out.WriteString(be.Right.ToString())
	out.WriteString(")")

	return out.String()
}

func (be *BinaryExpression) expressionNode() { /* EMPTY */ }
