package token

type Type string

const (
	ILLEGAL = "ILLEGAL"

	EOF = "EOF"

	IDENT = "IDENT"

	INT = "int"

	FLOAT = "float"

	STRING = "string"

	CHAR = "char"

	FUNCTION = "fn"

	LET = "let"

	VAR = "var"

	RETURN = "return"

	ASSIGN = "="

	PLUS = "+"

	MINUS = "-"

	MULTIPLY = "*"

	DIVIDE = "/"

	COMMA = ","

	SEMICOLON = ";"

	LPAREN = "("

	RPAREN = ")"

	LBRACE = "{"

	RBRACE = "}"

	NUMBER = "NUMBER"
)

var ReservedWords = map[string]Type{
	"let":    LET,
	"var":    VAR,
	"return": RETURN,
	"fn":     FUNCTION,
	"int":    INT,
	"float":  FLOAT,
	"string": STRING,
	"char":   CHAR,
}

type Token struct {
	Type    Type
	Literal string
}

func NewToken(t Type, lit string) *Token {
	return &Token{
		Type:    t,
		Literal: lit,
	}
}
