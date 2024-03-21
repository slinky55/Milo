package parser

import (
	"github.com/slinky55/milo/ast"
	"github.com/slinky55/milo/lexer"
	"github.com/slinky55/milo/token"
)

type Parser struct {
	lexer *lexer.Lexer
	curr  *token.Token
	next  *token.Token
}

type ParseError struct {
	What string
}

func (pe *ParseError) Error() string {
	return pe.What
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer: l,
	}

	p.curr = p.lexer.NextToken()
	p.next = p.lexer.NextToken()

	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}

	switch p.curr.Type {
	case token.LET:
		stmt, err := p.parseLetStmt()
		if err != nil {
			print(err.Error())
			break
		}
		program.Statements = append(program.Statements, stmt)
	}

	return program
}

func (p *Parser) advance() {
	p.curr = p.next
	p.next = p.lexer.NextToken()
}

func (p *Parser) parseLetStmt() (ast.Statement, error) {
	if p.next.Type != token.IDENT {
		return nil, &ParseError{
			What: string("Expected IDENT, found: " + p.next.Type),
		}
	}

	var stmt ast.LetStatement

	stmt.Token = p.curr

	p.advance()

	if p.next.Type != token.ASSIGN {
		return nil, &ParseError{
			What: string("Expected ASSIGN (\"=\"), found: " + p.next.Type),
		}
	}

	p.advance()
	p.advance()

	return &stmt, nil
}
