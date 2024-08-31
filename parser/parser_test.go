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

		println(stmt.ToString())
	}
}

func TestReturnStatement(t *testing.T) {
	input := "return 5; return foo; return 5 + foo;"
	l := lexer.New(input)
	p := New(l)

	program := p.Parse()

	for _, stmt := range program.Statements {
		if _, ok := stmt.(*ast.ReturnStatement); !ok {
			t.Error("not a return statement")
			continue
		}

		println(stmt.ToString())
	}
}

func TestPrefixExpressions(t *testing.T) {
	tests := []struct {
		Input    string
		Operator string
		Value    float64
	}{
		{"!5", "!", 5},
		{"--7;", "--", 7},
		{"++8", "++", 8},
		{"-1;", "-", 1},
	}

	for _, test := range tests {
		l := lexer.New(test.Input)
		p := New(l)

		program := p.Parse()

		if len(p.Errors) > 0 {
			t.Error("parser had errors")
			continue
		}

		if len(program.Statements) != 1 {
			t.Error("wrong number of statements")
			continue
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Error("not an expression statement")
			continue
		}

		expr, ok := stmt.Expr.(*ast.PrefixExpression)
		if !ok {
			t.Error("not a prefix expression")
			continue
		}

		if expr.Operator != test.Operator {
			t.Errorf("expected operator %s, found %s", test.Operator, expr.Operator)
		}

		num, ok := expr.Right.(*ast.NumberExpr)
		if !ok {
			t.Error("not a number expression")
			continue
		}

		if num.Value != test.Value {
			t.Errorf("expected value %f, found %f", test.Value, num.Value)
			continue
		}

		println(stmt.ToString())
	}
}

func TestBinaryExpressions(t *testing.T) {
	tests := []struct {
		input    string
		left     float64
		operator string
		right    float64
	}{
		{"5 + 5", 5, "+", 5},
		{"6 * 7;", 6, "*", 7},
		{"9 / 3", 9, "/", 3},
		{"2 + 2", 2, "+", 2},
		{"1 == 1", 1, "==", 1},
		{"0 != 2", 0, "!=", 2},
		{"10 > 6", 10, ">", 6},
		{"8 < 100", 8, "<", 100},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)

		program := p.Parse()

		if len(p.Errors) > 1 {
			t.Error("parser had errors")
			continue
		}

		if len(program.Statements) != 1 {
			t.Error("expected 1 statement, found ", len(program.Statements))
			continue
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Error("not an expression statement")
			continue
		}

		binary, ok := stmt.Expr.(*ast.BinaryExpression)
		if !ok {
			t.Error("not a binary expression")
			continue
		}

		left, ok := binary.Left.(*ast.NumberExpr)
		if !ok {
			t.Error("not a number expression")
			continue
		}

		if left.Value != test.left {
			t.Errorf("expected left value %f, found %f", test.left, left.Value)
			continue
		}

		if test.operator != binary.Operator {
			t.Errorf("expected operator %s, found %s", test.operator, binary.Operator)
		}

		right, ok := binary.Right.(*ast.NumberExpr)
		if !ok {
			t.Error("not a number expression")
			continue
		}

		if right.Value != test.right {
			t.Errorf("expected right value %f, found %f", test.right, right.Value)
			continue
		}

		println(stmt.ToString())
	}
}

func TestBooleanExpressions(t *testing.T) {
	tests := []struct {
		input    string
		left     bool
		operator string
		right    bool
	}{
		{"true == false;", true, "==", false},
		{"true != false", true, "!=", false},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.Parse()
		if len(p.Errors) > 0 {
			t.Error("parser had errors")
			continue
		}

		if len(program.Statements) != 1 {
			t.Error("wrong number of statements")
			continue
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Error("not an expression statement")
			continue
		}

		expr, ok := stmt.Expr.(*ast.BinaryExpression)
		if !ok {
			t.Error("not a binary expression")
			continue
		}

		left, ok := expr.Left.(*ast.BooleanExpr)
		if !ok {
			t.Error("not a boolean expression")
			continue
		}

		if left.Value != test.left {
			t.Errorf("expected left value %t, found %t", test.left, left.Value)
		}

		if expr.Operator != test.operator {
			t.Errorf("expected operator %s, found %s", test.operator, expr.Operator)
		}

		right, ok := expr.Right.(*ast.BooleanExpr)
		if !ok {
			t.Error("not a boolean expression")
			continue
		}

		if right.Value != test.right {
			t.Errorf("expected right value %t, found %t", test.right, right.Value)
			continue
		}

		println(stmt.ToString())
	}
}

func TestGroupedExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"(5 + 3) * 5", "((5 + 3) * 5)"},
		{"(5 + 3) * 5 * (5 + 3)", "(((5 + 3) * 5) * (5 + 3))"},
		{"(foo + bar) * bar / foo * (5 + 3)", "((((foo + bar) * bar) / foo) * (5 + 3))"},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)

		program := p.Parse()

		if len(p.Errors) > 0 {
			t.Error("parser had errors")
			continue
		}

		if len(program.Statements) != 1 {
			t.Error("wrong number of statements: ", len(program.Statements))
			continue
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Error("not an expression statement")
			continue
		}

		if stmt.ToString() != test.expected {
			t.Errorf("expected %s, found %s", test.expected, stmt.ToString())
			continue
		}

		println(stmt.ToString())
	}
}

func TestIfExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"if (5 * 3 == 1) { x } else { y }", "if (((5 * 3) == 1)) { x } else { y }"},
		{"if (true) { x }", "if (true) { x }"},
		{"let result = if (5 == 5) { true } else { false };", "let result = if ((5 == 5)) { true } else { false };"},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)

		program := p.Parse()

		if len(p.Errors) > 0 {
			t.Error("parser had errors")
			continue
		}

		if len(program.Statements) != 1 {
			t.Error("wrong number of statements: ", len(program.Statements))
			continue
		}

		stmt := program.Statements[0]

		if stmt.ToString() != test.expected {
			t.Errorf("expected %s, found %s", test.expected, stmt.ToString())
			continue
		}

		println(stmt.ToString())
	}
}

func TestFnExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"let result = fn (x, y) { return x + y; };", "let result = fn (x, y) { return (x + y); };"},
		{"return fn (x, y) { return x + y; };", "return fn (x, y) { return (x + y); };"},
		{"fn (foo, bar) { let x = foo; let y = bar; return x + y; }", "fn (foo, bar) { let x = foo;let y = bar;return (x + y); }"},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)

		program := p.Parse()

		if len(p.Errors) > 0 {
			t.Error("parser had errors")
			continue
		}

		if len(program.Statements) != 1 {
			t.Error("wrong number of statements: ", len(program.Statements))
			continue
		}

		stmt := program.Statements[0]

		if stmt.ToString() != test.expected {
			t.Errorf("expected %s, found %s", test.expected, stmt.ToString())
			continue
		}

		println(stmt.ToString())
	}
}

func TestCallExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"add(2, 3)", "add(2, 3)"},
		{"add(foo, bar)", "add(foo, bar)"},
		{"add(foo + 2, bar / 3)", "add((foo + 2), (bar / 3))"},
		{"add(foo, fn (x, y) { return x + y; })", "add(foo, fn (x, y) { return (x + y); })"},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)

		program := p.Parse()

		if len(p.Errors) > 0 {
			t.Error("parser had errors")
			continue
		}

		if len(program.Statements) != 1 {
			t.Error("wrong number of statements: ", len(program.Statements))
			continue
		}

		stmt := program.Statements[0]

		if stmt.ToString() != test.expected {
			t.Errorf("expected %s, found %s", test.expected, stmt.ToString())
			continue
		}

		println(stmt.ToString())
	}

}
