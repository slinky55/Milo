package parser

import (
	"github.com/slinky55/milo/ast"
	"github.com/slinky55/milo/token"
	"strconv"
)

func (p *Parser) parseExpr(precedence int) ast.Expression {
	var left ast.Expression
	switch p.cur.Type {
	case token.BANG, token.MINUS, token.INCREMENT, token.DECREMENT:
		left = p.parsePrefixExpr()
	case token.IDENT:
		left = p.parseIdentExpr()
	case token.NUMBER:
		left = p.parseNumberExpr()
	case token.TRUE, token.FALSE:
		left = p.parseBoolExpr()
	case token.IF:
		left = p.parseIfExpr()
	case token.FUNCTION:
		left = p.parseFunctionExpr()
	case token.LPAREN:
		left = p.parseGroupedExpression()
	default:
		p.error("unexpected %s at start of expression", p.cur.Literal)
		return nil
	}

	for (p.peek.Type != token.SEMICOLON) && (precedence < p.peekPrecedence()) {
		if !p.isBinaryOp(p.peek.Type) {
			return left
		}

		p.next()

		switch p.cur.Type {
		case token.LPAREN:
			left = p.parseCallExpr(left)
		default:
			left = p.parseBinaryExpression(left)
		}
	}

	return left
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

func (p *Parser) parseBoolExpr() *ast.BooleanExpr {
	return &ast.BooleanExpr{
		Token: p.cur,
		Value: p.cur.Type == token.TRUE,
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

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.next()

	expr := p.parseExpr(LOWEST)

	if !p.nextIfPeek(token.RPAREN) {
		return nil
	}

	return expr
}

func (p *Parser) parseIfExpr() *ast.IfExpr {
	expr := &ast.IfExpr{
		Token: p.cur,
	}

	if !p.nextIfPeek(token.LPAREN) {
		return nil
	}

	p.next()

	expr.Condition = p.parseExpr(LOWEST)

	if !p.nextIfPeek(token.RPAREN) {
		return nil
	}

	if !p.nextIfPeek(token.LBRACE) {
		return nil
	}

	expr.Consequence = p.parseStmtBlock()

	if p.peek.Type != token.ELSE {
		return expr
	}
	p.next()

	if !p.nextIfPeek(token.LBRACE) {
		return nil
	}

	expr.Alternative = p.parseStmtBlock()

	return expr
}

func (p *Parser) parseFunctionExpr() *ast.FunctionExpr {
	expr := &ast.FunctionExpr{
		Token: p.cur,
	}

	if !p.nextIfPeek(token.LPAREN) {
		return nil
	}

	expr.Parameters = p.parseParamList()

	if !p.nextIfPeek(token.LBRACE) {
		return nil
	}

	expr.Body = p.parseStmtBlock()

	return expr
}

func (p *Parser) parseParamList() []*ast.IdentExpr {
	var params []*ast.IdentExpr
	p.next()

	for {
		params = append(params, p.parseIdentExpr())

		if p.peek.Type == token.RPAREN {
			p.next()
			break
		}

		if !p.nextIfPeek(token.COMMA) {
			return nil
		}

		p.next()
	}

	return params
}

func (p *Parser) parseCallExpr(function ast.Expression) *ast.CallExpr {
	call := &ast.CallExpr{
		Token:    p.cur,
		Function: function,
	}

	call.Arguments = p.parseArgsList()
	return call
}

func (p *Parser) parseArgsList() []ast.Expression {
	var args []ast.Expression

	if p.peek.Type == token.RPAREN {
		p.next()
		return args
	}

	p.next()

	args = append(args, p.parseExpr(LOWEST))

	for p.peek.Type == token.COMMA {
		p.next()
		p.next()
		args = append(args, p.parseExpr(LOWEST))
	}

	if !p.nextIfPeek(token.RPAREN) {
		return nil
	}

	return args
}
