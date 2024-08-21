package lexer

import (
	"github.com/slinky55/milo/token"
	"testing"
)

func TestSingleCharTokens(t *testing.T) {
	input := "=;{}(),+-/*!<>"

	l := New(input)
	expected := []*token.Token{
		token.New(token.ASSIGN, "="),
		token.New(token.SEMICOLON, ";"),
		token.New(token.LBRACE, "{"),
		token.New(token.RBRACE, "}"),
		token.New(token.LPAREN, "("),
		token.New(token.RPAREN, ")"),
		token.New(token.COMMA, ","),
		token.New(token.PLUS, "+"),
		token.New(token.MINUS, "-"),
		token.New(token.DIVIDE, "/"),
		token.New(token.MULTIPLY, "*"),
		token.New(token.BANG, "!"),
		token.New(token.LTHAN, "<"),
		token.New(token.GTHAN, ">"),
	}

	for _, e := range expected {
		a := l.NextToken()

		if a.Type != e.Type {
			t.Errorf("expected type %s, found %s", e.Type, a.Type)
		}

		if a.Literal != e.Literal {
			t.Errorf("expected literal %s, found %s", e.Literal, a.Literal)
		}
	}

	last := l.NextToken()
	if last.Type != token.EOF {
		t.Errorf("expected EOF, found %s", last.Type)
	}
}

func TestTwoChar(t *testing.T) {
	input := "== !="

	l := New(input)
	expected := []*token.Token{
		token.New(token.EQUALS, "=="),
		token.New(token.NOTEQUALS, "!="),
	}

	for _, e := range expected {
		a := l.NextToken()

		if a.Type != e.Type {
			t.Errorf("expected type %s, found %s", e.Type, a.Type)
		}

		if a.Literal != e.Literal {
			t.Errorf("expected literal %s, found %s", e.Literal, a.Literal)
		}
	}

	last := l.NextToken()
	if last.Type != token.EOF {
		t.Errorf("expected EOF, found %s", last.Type)
	}
}

/*func TestComment(t *testing.T) {
	input := "let a = 5;//this is a comment\nvar b = 6;"

	l := New(input)

	expected := []*token.Token{
		token.New(token.LET, "let"),
		token.New(token.IDENT, "a"),
		token.New(token.ASSIGN, "="),
		token.New(token.NUMBER, "5"),
		token.New(token.SEMICOLON, ";"),
		token.New(token.VAR, "var"),
		token.New(token.IDENT, "b"),
		token.New(token.ASSIGN, "="),
		token.New(token.NUMBER, "6"),
		token.New(token.SEMICOLON, ";"),
	}

	for _, e := range expected {
		a := l.NextToken()

		if a.Type != e.Type {
			t.Errorf("expected type %s, found %s", e.Type, a.Type)
		}

		if a.Literal != e.Literal {
			t.Errorf("expected literal %s, found %s", e.Literal, a.Literal)
		}
	}

	last := l.NextToken()
	if last.Type != token.EOF {
		t.Errorf("expected EOF, found %s", last.Type)
	}
}*/

func TestReservedWords(t *testing.T) {
	input := "let var return fn int float string char true false"
	l := New(input)

	expected := []*token.Token{
		token.New(token.LET, "let"),
		token.New(token.VAR, "var"),
		token.New(token.RETURN, "return"),
		token.New(token.FUNCTION, "fn"),
		token.New(token.INT, "int"),
		token.New(token.FLOAT, "float"),
		token.New(token.STRING, "string"),
		token.New(token.CHAR, "char"),
		token.New(token.TRUE, "true"),
		token.New(token.FALSE, "false"),
	}

	for _, e := range expected {
		a := l.NextToken()

		if a.Type != e.Type {
			t.Errorf("expected type %s, found %s", e.Type, a.Type)
		}

		if a.Literal != e.Literal {
			t.Errorf("expected literal %s, found %s", e.Literal, a.Literal)
		}
	}

	last := l.NextToken()
	if last.Type != token.EOF {
		t.Errorf("expected EOF, found %s", last.Type)
	}
}

func TestIdentAndNumber(t *testing.T) {
	input := "let a = 5; var foo = 600;"
	l := New(input)

	expected := []*token.Token{
		token.New(token.LET, "let"),
		token.New(token.IDENT, "a"),
		token.New(token.ASSIGN, "="),
		token.New(token.NUMBER, "5"),
		token.New(token.SEMICOLON, ";"),
		token.New(token.VAR, "var"),
		token.New(token.IDENT, "foo"),
		token.New(token.ASSIGN, "="),
		token.New(token.NUMBER, "600"),
		token.New(token.SEMICOLON, ";"),
	}

	for _, e := range expected {
		a := l.NextToken()

		if a.Type != e.Type {
			t.Errorf("expected type %s, found %s", e.Type, a.Type)
		}

		if a.Literal != e.Literal {
			t.Errorf("expected literal %s, found %s", e.Literal, a.Literal)
		}
	}

	last := l.NextToken()
	if last.Type != token.EOF {
		t.Errorf("expected EOF, found %s", last.Type)
	}
}

func TestMath(t *testing.T) {
	input := "let a = 5 + 6; let b = 7 / 3; let c = 6 * 6; let d = 5 / 5; let foo = a * b; let bar = c + d; let baz = foo / bar;"

	l := New(input)

	expected := []*token.Token{
		token.New(token.LET, "let"),
		token.New(token.IDENT, "a"),
		token.New(token.ASSIGN, "="),
		token.New(token.NUMBER, "5"),
		token.New(token.PLUS, "+"),
		token.New(token.NUMBER, "6"),
		token.New(token.SEMICOLON, ";"),

		token.New(token.LET, "let"),
		token.New(token.IDENT, "b"),
		token.New(token.ASSIGN, "="),
		token.New(token.NUMBER, "7"),
		token.New(token.DIVIDE, "/"),
		token.New(token.NUMBER, "3"),
		token.New(token.SEMICOLON, ";"),

		token.New(token.LET, "let"),
		token.New(token.IDENT, "c"),
		token.New(token.ASSIGN, "="),
		token.New(token.NUMBER, "6"),
		token.New(token.MULTIPLY, "*"),
		token.New(token.NUMBER, "6"),
		token.New(token.SEMICOLON, ";"),

		token.New(token.LET, "let"),
		token.New(token.IDENT, "d"),
		token.New(token.ASSIGN, "="),
		token.New(token.NUMBER, "5"),
		token.New(token.DIVIDE, "/"),
		token.New(token.NUMBER, "5"),
		token.New(token.SEMICOLON, ";"),

		token.New(token.LET, "let"),
		token.New(token.IDENT, "foo"),
		token.New(token.ASSIGN, "="),
		token.New(token.IDENT, "a"),
		token.New(token.MULTIPLY, "*"),
		token.New(token.IDENT, "b"),
		token.New(token.SEMICOLON, ";"),

		token.New(token.LET, "let"),
		token.New(token.IDENT, "bar"),
		token.New(token.ASSIGN, "="),
		token.New(token.IDENT, "c"),
		token.New(token.PLUS, "+"),
		token.New(token.IDENT, "d"),
		token.New(token.SEMICOLON, ";"),

		token.New(token.LET, "let"),
		token.New(token.IDENT, "baz"),
		token.New(token.ASSIGN, "="),
		token.New(token.IDENT, "foo"),
		token.New(token.DIVIDE, "/"),
		token.New(token.IDENT, "bar"),
		token.New(token.SEMICOLON, ";"),
	}

	for _, e := range expected {
		a := l.NextToken()

		if a.Type != e.Type {
			t.Errorf("expected type %s, found %s", e.Type, a.Type)
		}

		if a.Literal != e.Literal {
			t.Errorf("expected literal %s, found %s", e.Literal, a.Literal)
		}
	}

	last := l.NextToken()
	if last.Type != token.EOF {
		t.Errorf("expected EOF, found %s", last.Type)
	}
}

func TestFunctionDefinition(t *testing.T) {
	input := "fn foo(int x, int y) { return x + y; }"

	l := New(input)

	expected := []*token.Token{
		token.New(token.FUNCTION, "fn"),
		token.New(token.IDENT, "foo"),
		token.New(token.LPAREN, "("),
		token.New(token.INT, "int"),
		token.New(token.IDENT, "x"),
		token.New(token.COMMA, ","),
		token.New(token.INT, "int"),
		token.New(token.IDENT, "y"),
		token.New(token.RPAREN, ")"),
		token.New(token.LBRACE, "{"),
		token.New(token.RETURN, "return"),
		token.New(token.IDENT, "x"),
		token.New(token.PLUS, "+"),
		token.New(token.IDENT, "y"),
		token.New(token.SEMICOLON, ";"),
		token.New(token.RBRACE, "}"),
	}

	for _, e := range expected {
		a := l.NextToken()

		if a.Type != e.Type {
			t.Errorf("expected type %s, found %s", e.Type, a.Type)
		}

		if a.Literal != e.Literal {
			t.Errorf("expected literal %s, found %s", e.Literal, a.Literal)
		}
	}

	last := l.NextToken()
	if last.Type != token.EOF {
		t.Errorf("expected EOF, found %s", last.Type)
	}
}
