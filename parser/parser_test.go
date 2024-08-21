package parser

import (
	"fmt"
	"github.com/slinky55/milo/ast"
	"github.com/slinky55/milo/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := "let a = 5; let b = 6; let c = 7;"
	l := lexer.New(input)
	p := New(l)

	program := p.Parse()

	idents := []string{
		"a", "b", "c",
	}

	values := []float64{
		5, 6, 7,
	}

	for i, stmt := range program.Statements {
		let, ok := stmt.(*ast.LetStatement)
		if !ok {
			p.error("not a let statement")
			continue
		}

		if let.Ident.Literal() != idents[i] {
			p.error(fmt.Sprintf("expected ident %s, found %s", idents[i], let.Ident.Literal()))
		}

		num, ok := let.Expr.(*ast.NumberExpr)
		if !ok {
			p.error("not a number expression")
			continue
		}

		if num.Value != values[i] {
			p.error(fmt.Sprintf("expected value %f, found %f", values[i], num.Value))
		}
	}
}
