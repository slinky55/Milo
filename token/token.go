package token

type Type string

const (
	ILLEGAL = "ILLEGAL"

	EOF = "EOF"

  NUMBER = "NUMBER"

	IDENT = "IDENT"

	INT = "INT"

	ASSIGN = ":="

	PLUS = "+"

	COMMA = ","

	SEMICOLON = ";"

	LPAREN = "("

	RPAREN = ")"

	LBRACE = "{"

	RBRACE = "}"

	FUNCTION = "fn"

	LET = "let"
)

type Token struct {
	Type    Type
	Literal string
}

