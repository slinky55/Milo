package parser

import (
	"fmt"
	"github.com/slinky55/milo/ast"
	"github.com/slinky55/milo/lexer"
	"github.com/slinky55/milo/token"
	"strconv"
)

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
		p.next()
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
	expr := p.parseExpr()
	if expr == nil {
		return nil
	}

	p.next()
	if p.cur.Type != token.SEMICOLON {
		p.error(fmt.Sprintf("expected %s, but found %s", token.SEMICOLON, p.cur.Type))
		return nil
	}

	return &ast.LetStatement{
		Token: t,
		Ident: ident,
		Expr:  expr,
	}
}

func (p *Parser) parseReturnStmt() *ast.ReturnStatement {
	t := p.cur
	p.next()
	expr := p.parseExpr()
	if expr == nil {
		return nil
	}
	return &ast.ReturnStatement{
		Token: t,
		Expr:  expr,
	}
}

func (p *Parser) parseExpr() ast.Expression {
	t := p.cur
	switch t.Type {
	case token.NUMBER:
		return p.parseNumberExpr()
	case token.IDENT:
		return p.parseIdentExpr()
	default:
		p.error("unknown expression")
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

func (p *Parser) next() {
	p.cur = p.peek
	p.peek = p.l.NextToken()
}

func (p *Parser) error(msg string) {
	fmt.Printf("parser error: %s\n", msg)
}
