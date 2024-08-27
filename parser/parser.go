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

type Parser struct {
	l *lexer.Lexer

	cur  *token.Token
	peek *token.Token
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

		// TODO: Consider adding expression statements
		switch p.cur.Type {
		case token.LET:
			stmt = p.parseLetStmt()
		case token.RETURN:
			stmt = p.parseReturnStmt()
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

func (p *Parser) parseExpr(precedence int) ast.Expression {
	switch p.cur.Type {
	case token.BANG, token.MINUS, token.INCREMENT, token.DECREMENT:
		return p.parsePrefixExpr()
	case token.NUMBER:
		return p.parseNumberExpr()
	case token.IDENT:
		return p.parseIdentExpr()
	default:
		p.error("unknown start of expression")
		return nil
	}
}

func (p *Parser) parseIdentExpr() *ast.IdentExpr {
	return &ast.IdentExpr{
		Token: p.cur,
		Value: p.cur.Literal,
	}
}

func (p *Parser) parseNumberExpr() *ast.NumberExpr {
	value, err := strconv.ParseFloat(p.cur.Literal, 64)
	if err != nil {
		p.error(err.Error())
		return nil
	}
	return &ast.NumberExpr{
		Token: p.cur,
		Value: value,
	}
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

func (p *Parser) error(msg string) {
	fmt.Printf("parser error: %s\n", msg)
}
