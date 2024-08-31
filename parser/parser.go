package parser

import (
	"fmt"
	"github.com/slinky55/milo/ast"
	"github.com/slinky55/milo/lexer"
	"github.com/slinky55/milo/token"
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
	token.LPAREN:    CALL,
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
	token.LPAREN:    "",
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
		stmt := p.parseStatement()

		if stmt != nil {
			program.AddStatement(stmt)
		} else {
			break
		}
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	var stmt ast.Statement

	switch p.cur.Type {
	case token.LET:
		stmt = p.parseLetStmt()
	case token.RETURN:
		stmt = p.parseReturnStmt()
	default:
		stmt = p.parseExprStatement()
	}

	return stmt
}

func (p *Parser) parseLetStmt() *ast.LetStatement {
	t := p.cur

	if !p.nextIfPeek(token.IDENT) {
		return nil
	}

	ident := p.parseIdentExpr()

	if !p.nextIfPeek(token.ASSIGN) {
		return nil
	}

	p.next()

	expr := p.parseExpr(LOWEST)

	if expr == nil {
		return nil
	}

	if !p.nextIfPeek(token.SEMICOLON) {
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
		p.error("expected %s, but found %s", token.SEMICOLON, p.cur.Type)
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

func (p *Parser) parseStmtBlock() *ast.StatementBlock {
	block := &ast.StatementBlock{
		Token: p.cur,
	}

	p.next()

	for p.cur.Type != token.RBRACE && p.cur.Type != token.EOF {
		stmt := p.parseStatement()

		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
	}

	return block
}

func (p *Parser) next() {
	p.cur = p.peek
	p.peek = p.l.NextToken()
}

func (p *Parser) nextIfPeek(t token.Type) bool {
	if p.peek.Type != t {
		p.error("expected %s, but found %s", ")", p.peek.Literal)
		return false
	}
	p.next()
	return true
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

func (p *Parser) error(msg string, args ...any) {
	err := "parser error: " + fmt.Sprintf(msg, args...)
	p.Errors = append(p.Errors, err)
	//println(err)
}
