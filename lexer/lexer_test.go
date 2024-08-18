package lexer

import (
	"github.com/slinky55/milo/token"
	"testing"
)

func TestSingleCharTokens(t *testing.T) {
	input := "=;{}(),+-/*"

	l := NewLexer(input)
	expected := []*token.Token{
		token.NewToken(token.ASSIGN, "="),
		token.NewToken(token.SEMICOLON, ";"),
		token.NewToken(token.LBRACE, "{"),
		token.NewToken(token.RBRACE, "}"),
		token.NewToken(token.LPAREN, "("),
		token.NewToken(token.RPAREN, ")"),
		token.NewToken(token.COMMA, ","),
		token.NewToken(token.PLUS, "+"),
		token.NewToken(token.MINUS, "-"),
		token.NewToken(token.DIVIDE, "/"),
		token.NewToken(token.MULTIPLY, "*"),
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

func TestReservedWords(t *testing.T) {
	input := "let var return fn int float string char"
	l := NewLexer(input)

	expected := []*token.Token{
		token.NewToken(token.LET, "let"),
		token.NewToken(token.VAR, "var"),
		token.NewToken(token.RETURN, "return"),
		token.NewToken(token.FUNCTION, "fn"),
		token.NewToken(token.INT, "int"),
		token.NewToken(token.FLOAT, "float"),
		token.NewToken(token.STRING, "string"),
		token.NewToken(token.CHAR, "char"),
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
	l := NewLexer(input)

	expected := []*token.Token{
		token.NewToken(token.LET, "let"),
		token.NewToken(token.IDENT, "a"),
		token.NewToken(token.ASSIGN, "="),
		token.NewToken(token.NUMBER, "5"),
		token.NewToken(token.SEMICOLON, ";"),
		token.NewToken(token.VAR, "var"),
		token.NewToken(token.IDENT, "foo"),
		token.NewToken(token.ASSIGN, "="),
		token.NewToken(token.NUMBER, "600"),
		token.NewToken(token.SEMICOLON, ";"),
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

	l := NewLexer(input)

	expected := []*token.Token{
		token.NewToken(token.LET, "let"),
		token.NewToken(token.IDENT, "a"),
		token.NewToken(token.ASSIGN, "="),
		token.NewToken(token.NUMBER, "5"),
		token.NewToken(token.PLUS, "+"),
		token.NewToken(token.NUMBER, "6"),
		token.NewToken(token.SEMICOLON, ";"),

		token.NewToken(token.LET, "let"),
		token.NewToken(token.IDENT, "b"),
		token.NewToken(token.ASSIGN, "="),
		token.NewToken(token.NUMBER, "7"),
		token.NewToken(token.DIVIDE, "/"),
		token.NewToken(token.NUMBER, "3"),
		token.NewToken(token.SEMICOLON, ";"),

		token.NewToken(token.LET, "let"),
		token.NewToken(token.IDENT, "c"),
		token.NewToken(token.ASSIGN, "="),
		token.NewToken(token.NUMBER, "6"),
		token.NewToken(token.MULTIPLY, "*"),
		token.NewToken(token.NUMBER, "6"),
		token.NewToken(token.SEMICOLON, ";"),

		token.NewToken(token.LET, "let"),
		token.NewToken(token.IDENT, "d"),
		token.NewToken(token.ASSIGN, "="),
		token.NewToken(token.NUMBER, "5"),
		token.NewToken(token.DIVIDE, "/"),
		token.NewToken(token.NUMBER, "5"),
		token.NewToken(token.SEMICOLON, ";"),

		token.NewToken(token.LET, "let"),
		token.NewToken(token.IDENT, "foo"),
		token.NewToken(token.ASSIGN, "="),
		token.NewToken(token.IDENT, "a"),
		token.NewToken(token.MULTIPLY, "*"),
		token.NewToken(token.IDENT, "b"),
		token.NewToken(token.SEMICOLON, ";"),

		token.NewToken(token.LET, "let"),
		token.NewToken(token.IDENT, "bar"),
		token.NewToken(token.ASSIGN, "="),
		token.NewToken(token.IDENT, "c"),
		token.NewToken(token.PLUS, "+"),
		token.NewToken(token.IDENT, "d"),
		token.NewToken(token.SEMICOLON, ";"),

		token.NewToken(token.LET, "let"),
		token.NewToken(token.IDENT, "baz"),
		token.NewToken(token.ASSIGN, "="),
		token.NewToken(token.IDENT, "foo"),
		token.NewToken(token.DIVIDE, "/"),
		token.NewToken(token.IDENT, "bar"),
		token.NewToken(token.SEMICOLON, ";"),
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

	l := NewLexer(input)

	expected := []*token.Token{
		token.NewToken(token.FUNCTION, "fn"),
		token.NewToken(token.IDENT, "foo"),
		token.NewToken(token.LPAREN, "("),
		token.NewToken(token.INT, "int"),
		token.NewToken(token.IDENT, "x"),
		token.NewToken(token.COMMA, ","),
		token.NewToken(token.INT, "int"),
		token.NewToken(token.IDENT, "y"),
		token.NewToken(token.RPAREN, ")"),
		token.NewToken(token.LBRACE, "{"),
		token.NewToken(token.RETURN, "return"),
		token.NewToken(token.IDENT, "x"),
		token.NewToken(token.PLUS, "+"),
		token.NewToken(token.IDENT, "y"),
		token.NewToken(token.SEMICOLON, ";"),
		token.NewToken(token.RBRACE, "}"),
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
