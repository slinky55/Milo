package ast

import (
	"github.com/slinky55/milo/token"
	"strings"
)

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

func (ls *LetStatement) ToString() string {
	var out strings.Builder

	out.WriteString(ls.Literal())
	out.WriteString(" " + ls.Ident.ToString() + " ")
	out.WriteString("= ")

	if ls.Expr != nil {
		out.WriteString(ls.Expr.ToString())
	}

	out.WriteString(";")

	return out.String()
}

func (ls *LetStatement) statementNode() { /* EMPTY */ }

type ReturnStatement struct {
	Token *token.Token
	Expr  Expression
}

func (rs *ReturnStatement) Literal() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) ToString() string {
	var out strings.Builder

	out.WriteString(rs.Literal())
	out.WriteString(" ")
	out.WriteString(rs.Expr.ToString())

	out.WriteString(";")

	return out.String()
}

func (rs *ReturnStatement) statementNode() { /* EMPTY */ }

type ExpressionStatement struct {
	Token *token.Token
	Expr  Expression
}

func (es *ExpressionStatement) Literal() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) ToString() string {
	return es.Expr.ToString()
}

func (es *ExpressionStatement) statementNode() { /* EMPTY */ }

type StatementBlock struct {
	Token      *token.Token
	Statements []Statement
}

func (es *StatementBlock) Literal() string {
	return es.Token.Literal
}

func (es *StatementBlock) ToString() string {
	var out strings.Builder

	out.WriteString("{ ")
	for _, stmt := range es.Statements {
		out.WriteString(stmt.ToString())
	}
	out.WriteString(" }")

	return out.String()
}

func (es *StatementBlock) statementNode() { /* EMPTY */ }
