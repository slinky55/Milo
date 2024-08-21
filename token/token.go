package token

type Type string

const (
	ILLEGAL = "ILLEGAL"

	EOF = "EOF"

	IDENT = "IDENT"

	NUMBER = "NUMBER"

	INT = "INT"

	FLOAT = "FLOAT"

	STRING = "STRING"

	CHAR = "CHAR"

	FUNCTION = "FUNCTION"

	LET = "LET"

	VAR = "VAR"

	RETURN = "RETURN"

	TRUE = "TRUE"

	FALSE = "FALSE"

	ASSIGN = "ASSIGN"

	PLUS = "PLUS"

	MINUS = "MINUS"

	MULTIPLY = "MULTIPLY"

	DIVIDE = "DIVIDE"

	COMMA = "COMMA"

	SEMICOLON = "SEMICOLON"

	LPAREN = "LPAREN"

	RPAREN = "RPAREN"

	LBRACE = "LBRACE"

	RBRACE = "RBRACE"

	BANG = "BANG"

	LTHAN = "LESS THEN"

	GTHAN = "GREATER THEN"

	EQUALS = "EQUALS"

	NOTEQUALS = "NOT EQUALS"
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
	"true":   TRUE,
	"false":  FALSE,
}

type Token struct {
	Type    Type
	Literal string
}

func New(t Type, lit string) *Token {
	return &Token{
		Type:    t,
		Literal: lit,
	}
}
