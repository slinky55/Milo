package lexer

import (
	"github.com/slinky55/milo/token"
	"unicode"
)

type Lexer struct {
	input   string
	charPos int
	readPos int
	char    byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input:   input,
		charPos: 0,
		readPos: 0,
		char:    0,
	}
	l.advance()
	return l
}

func (l *Lexer) NextToken() *token.Token {
	var tk *token.Token

	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.advance()
	}

	switch l.char {
	case '=':
		tk = token.NewToken(token.ASSIGN, string(l.char))
	case ';':
		tk = token.NewToken(token.SEMICOLON, string(l.char))
	case ',':
		tk = token.NewToken(token.COMMA, string(l.char))
	case '{':
		tk = token.NewToken(token.LBRACE, string(l.char))
	case '}':
		tk = token.NewToken(token.RBRACE, string(l.char))
	case '(':
		tk = token.NewToken(token.LPAREN, string(l.char))
	case ')':
		tk = token.NewToken(token.RPAREN, string(l.char))
	case '+':
		tk = token.NewToken(token.PLUS, string(l.char))
	case '-':
		tk = token.NewToken(token.MINUS, string(l.char))
	case '*':
		tk = token.NewToken(token.MULTIPLY, string(l.char))
	case '/':
		tk = token.NewToken(token.DIVIDE, string(l.char))
	case 0:
		tk = token.NewToken(token.EOF, "")
	default:
		if unicode.IsLetter(rune(l.char)) {
			start := l.charPos
			for unicode.IsLetter(rune(l.char)) || l.char == '_' {
				l.advance()
			}
			literal := l.input[start:l.charPos]

			if t, ok := token.ReservedWords[literal]; ok {
				tk = token.NewToken(t, literal)
			} else {
				tk = token.NewToken(token.IDENT, literal)
			}
		} else if unicode.IsNumber(rune(l.char)) {
			start := l.charPos
			for unicode.IsNumber(rune(l.char)) {
				l.advance()
			}
			literal := l.input[start:l.charPos]

			tk = token.NewToken(token.NUMBER, literal)
		} else {
			tk = token.NewToken(token.ILLEGAL, string(l.char))
		}
		return tk
	}

	l.advance()
	return tk
}

func (l *Lexer) advance() {
	if l.readPos >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPos]
	}
	l.charPos = l.readPos
	l.readPos++
}
