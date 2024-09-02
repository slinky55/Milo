package token

type Type string

const (
	ILLEGAL = "ILLEGAL"

	EOF = "EOF"

	IDENT = "IDENT"

	NUMBER = "NUMBER"

	STRING = "STRING"

	FUNCTION = "FUNCTION"

	LET = "LET"

	VAR = "VAR"

	RETURN = "RETURN"

	TRUE = "TRUE"

	FALSE = "FALSE"

	IF = "IF"

	ELSE = "ELSE"

	NULL = "NULL"

	ASSIGN = "ASSIGN"

	PLUS = "PLUS"

	MINUS = "MINUS"

	MULTIPLY = "MULTIPLY"

	DIVIDE = "DIVIDE"

	INCREMENT = "INCREMENT"

	DECREMENT = "DECREMENT"

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
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"null":   NULL,
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
