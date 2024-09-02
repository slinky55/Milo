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

type StringExpr struct {
	Token *token.Token
	Value string
}

func (se *StringExpr) Literal() string {
	return se.Token.Literal
}

func (se *StringExpr) ToString() string {
	return se.Literal()
}

func (se *StringExpr) expressionNode() { /* EMPTY */ }

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

type BooleanExpr struct {
	Token *token.Token
	Value bool
}

func (be *BooleanExpr) Literal() string {
	return be.Token.Literal
}

func (be *BooleanExpr) ToString() string {
	return be.Literal()
}

func (be *BooleanExpr) expressionNode() { /* EMPTY */ }

type IfExpr struct {
	Token       *token.Token
	Condition   Expression
	Consequence *StatementBlock
	Alternative *StatementBlock
}

func (ie *IfExpr) Literal() string {
	return ie.Token.Literal
}

func (ie *IfExpr) ToString() string {
	var out strings.Builder
	out.WriteString("if")
	out.WriteString(" (" + ie.Condition.ToString() + ") ")
	out.WriteString(ie.Consequence.ToString())

	if ie.Alternative != nil {
		out.WriteString(" else " + ie.Alternative.ToString())
	}

	return out.String()
}

func (ie *IfExpr) expressionNode() { /* EMPTY */ }

type FunctionExpr struct {
	Token      *token.Token
	Parameters []*IdentExpr
	Body       *StatementBlock
}

func (fe *FunctionExpr) Literal() string {
	return fe.Token.Literal
}

func (fe *FunctionExpr) ToString() string {
	var out strings.Builder

	var params []string
	for _, p := range fe.Parameters {
		params = append(params, p.ToString())
	}

	out.WriteString(fe.Literal())
	out.WriteString(" (")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fe.Body.ToString())

	return out.String()
}

func (fe *FunctionExpr) expressionNode() { /* EMPTY */ }

type CallExpr struct {
	Token     *token.Token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpr) Literal() string {
	return ce.Token.Literal
}

func (ce *CallExpr) ToString() string {
	var out strings.Builder
	var args []string

	for _, a := range ce.Arguments {
		args = append(args, a.ToString())
	}

	out.WriteString(ce.Function.ToString())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

func (ce *CallExpr) expressionNode() { /* EMPTY */ }
