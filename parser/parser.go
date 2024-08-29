package parser

import (
	"fmt"
	"github.com/slinky55/milo/ast"
	"github.com/slinky55/milo/lexer"
	"github.com/slinky55/milo/token"
	"strconv"
)

// Operator precedence
const (
	_ int = iota
	LOWEST
	EQUALITY
	COMPARISON
	SUM
	PRODUCT
	PREFIX
	CALL
)

var TokenPrecedence = map[token.Type]int{
	token.EQUALS:    EQUALITY,
	token.NOTEQUALS: EQUALITY,
	token.LTHAN:     COMPARISON,
	token.GTHAN:     COMPARISON,
	token.PLUS:      SUM,
	token.MINUS:     SUM,
	token.MULTIPLY:  PRODUCT,
	token.DIVIDE:    PRODUCT,
}

var BinaryOps = map[token.Type]string{
	token.MULTIPLY:  "",
	token.DIVIDE:    "",
	token.PLUS:      "",
	token.MINUS:     "",
	token.EQUALS:    "",
	token.NOTEQUALS: "",
	token.GTHAN:     "",
	token.LTHAN:     "",
}

var PrefixOps = map[token.Type]string{
	token.BANG:      "",
	token.INCREMENT: "",
	token.DECREMENT: "",
}

type Parser struct {
	l *lexer.Lexer

	cur  *token.Token
	peek *token.Token

	Errors []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.cur = p.l.NextToken()
	p.peek = p.l.NextToken()

	return p
}

func (p *Parser) Parse() *ast.Program {
	program := &ast.Program{}

	for p.cur.Type != token.EOF {
		var stmt ast.Statement

		switch p.cur.Type {
		case token.LET:
			stmt = p.parseLetStmt()
		case token.RETURN:
			stmt = p.parseReturnStmt()
		default:
			stmt = p.parseExprStatement()
		}

		if stmt != nil {
			program.AddStatement(stmt)
		} else {
			break
		}
	}

	return program
}

func (p *Parser) parseLetStmt() *ast.LetStatement {
	t := p.cur

	p.next()
	if p.cur.Type != token.IDENT {
		p.error(fmt.Sprintf("expected %s, but found %s", token.IDENT, p.cur.Type))
		return nil
	}
	ident := p.parseIdentExpr()

	p.next()
	if p.cur.Type != token.ASSIGN {
		p.error(fmt.Sprintf("expected %s, but found %s", token.ASSIGN, p.cur.Type))
		return nil
	}

	p.next()
	expr := p.parseExpr(LOWEST)
	if expr == nil {
		return nil
	}

	p.next()
	if p.cur.Type != token.SEMICOLON {
		p.error(fmt.Sprintf("expected %s, but found %s", token.SEMICOLON, p.cur.Type))
		return nil
	}
	p.next()

	return &ast.LetStatement{
		Token: t,
		Ident: ident,
		Expr:  expr,
	}
}

func (p *Parser) parseReturnStmt() *ast.ReturnStatement {
	t := p.cur
	p.next()
	expr := p.parseExpr(LOWEST)
	if expr == nil {
		return nil
	}

	p.next()
	if p.cur.Type != token.SEMICOLON {
		p.error(fmt.Sprintf("expected %s, but found %s", token.SEMICOLON, p.cur.Type))
		return nil
	}
	p.next()

	return &ast.ReturnStatement{
		Token: t,
		Expr:  expr,
	}
}

func (p *Parser) parseExprStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{
		Token: p.cur,
	}

	stmt.Expr = p.parseExpr(LOWEST)
	if stmt.Expr == nil {
		return nil
	}

	if p.peek.Type == token.SEMICOLON {
		p.next()
	}

	p.next()

	return stmt
}

func (p *Parser) parseExpr(precedence int) ast.Expression {
	var left ast.Expression
	switch p.cur.Type {
	case token.BANG, token.MINUS, token.INCREMENT, token.DECREMENT:
		left = p.parsePrefixExpr()
	case token.IDENT:
		left = p.parseIdentExpr()
	case token.NUMBER:
		left = p.parseNumberExpr()
	default:
		p.error(fmt.Sprintf("unexpected %s at start of expression", p.cur.Literal))
		return nil
	}

	for (p.peek.Type != token.SEMICOLON) && (precedence < p.peekPrecedence()) {
		if !p.isBinaryOp(p.peek.Type) {
			return left
		}

		p.next()

		left = p.parseBinaryExpression(left)
	}

	return left
}

func (p *Parser) parseIdentExpr() *ast.IdentExpr {
	expr := &ast.IdentExpr{
		Token: p.cur,
		Value: p.cur.Literal,
	}
	return expr
}

func (p *Parser) parseNumberExpr() *ast.NumberExpr {
	value, err := strconv.ParseFloat(p.cur.Literal, 64)
	if err != nil {
		p.error(err.Error())
		return nil
	}
	expr := &ast.NumberExpr{
		Token: p.cur,
		Value: value,
	}
	return expr
}

func (p *Parser) parsePrefixExpr() *ast.PrefixExpression {
	expr := &ast.PrefixExpression{
		Token:    p.cur,
		Operator: p.cur.Literal,
	}

	p.next()

	expr.Right = p.parseExpr(PREFIX)

	return expr
}

func (p *Parser) parseBinaryExpression(left ast.Expression) *ast.BinaryExpression {
	expr := &ast.BinaryExpression{
		Token:    p.cur,
		Operator: p.cur.Literal,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.next()
	expr.Right = p.parseExpr(precedence)

	return expr
}

func (p *Parser) next() {
	p.cur = p.peek
	p.peek = p.l.NextToken()
}

func (p *Parser) curPrecedence() int {
	if pr, ok := TokenPrecedence[p.cur.Type]; ok {
		return pr
	}
	return LOWEST
}

func (p *Parser) peekPrecedence() int {
	if pr, ok := TokenPrecedence[p.peek.Type]; ok {
		return pr
	}
	return LOWEST
}

func (p *Parser) isBinaryOp(t token.Type) bool {
	if _, ok := BinaryOps[t]; ok {
		return true
	}
	return false
}

func (p *Parser) isPrefixOp(t token.Type) bool {
	if _, ok := PrefixOps[t]; ok {
		return true
	}
	return false
}

func (p *Parser) error(msg string) {
	err := fmt.Sprintf("parser error: %s\n", msg)
	p.Errors = append(p.Errors, err)
	println(err)
}
