package token

type Type string

const (
	ILLEGAL = "ILLEGAL"

	EOF = "EOF"

	NUMBER = "NUMBER"

	IDENT = "IDENT"

	INT = "int"

	STRING = "string"

	ASSIGN = "="

	PLUS = "+"

	COMMA = ","

	SEMICOLON = ";"

	LPAREN = "("

	RPAREN = ")"

	LBRACE = "{"

	RBRACE = "}"

	FUNCTION = "fn"

	LET = "let"

	CONST = "const"

	RETURN = "return"
)

type Token struct {
	Type    Type
	Literal string
}
