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

func New(input string) *Lexer {
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
		if l.peek() == '=' {
			first := string(l.char)
			l.advance()
			literal := first + string(l.char)
			tk = token.New(token.EQUALS, literal)
		} else {
			tk = token.New(token.ASSIGN, string(l.char))
		}
	case ';':
		tk = token.New(token.SEMICOLON, string(l.char))
	case ',':
		tk = token.New(token.COMMA, string(l.char))
	case '{':
		tk = token.New(token.LBRACE, string(l.char))
	case '}':
		tk = token.New(token.RBRACE, string(l.char))
	case '(':
		tk = token.New(token.LPAREN, string(l.char))
	case ')':
		tk = token.New(token.RPAREN, string(l.char))
	case '+':
		if l.peek() == '+' {
			first := string(l.char)
			l.advance()
			literal := first + string(l.char)
			tk = token.New(token.INCREMENT, literal)
		} else {
			tk = token.New(token.PLUS, string(l.char))
		}
	case '-':
		if l.peek() == '-' {
			first := string(l.char)
			l.advance()
			literal := first + string(l.char)
			tk = token.New(token.DECREMENT, literal)
		} else {
			tk = token.New(token.MINUS, string(l.char))
		}
	case '*':
		tk = token.New(token.MULTIPLY, string(l.char))
	case '/':
		if l.peek() == '/' {
			l.advance()
			for l.peek() != '\n' {
				l.advance()
			}
		} else {
			tk = token.New(token.DIVIDE, string(l.char))
		}
	case '!':
		if l.peek() == '=' {
			first := string(l.char)
			l.advance()
			literal := first + string(l.char)
			tk = token.New(token.NOTEQUALS, literal)
		} else {
			tk = token.New(token.BANG, string(l.char))
		}
	case '<':
		tk = token.New(token.LTHAN, string(l.char))
	case '>':
		tk = token.New(token.GTHAN, string(l.char))
	case 0:
		tk = token.New(token.EOF, "")
	default:
		if unicode.IsLetter(rune(l.char)) {
			start := l.charPos
			for unicode.IsLetter(rune(l.char)) || l.char == '_' {
				l.advance()
			}
			literal := l.input[start:l.charPos]

			if t, ok := token.ReservedWords[literal]; ok {
				tk = token.New(t, literal)
			} else {
				tk = token.New(token.IDENT, literal)
			}
		} else if unicode.IsNumber(rune(l.char)) {
			start := l.charPos
			for unicode.IsNumber(rune(l.char)) {
				l.advance()
			}
			literal := l.input[start:l.charPos]

			tk = token.New(token.NUMBER, literal)
		} else {
			tk = token.New(token.ILLEGAL, string(l.char))
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

func (l *Lexer) peek() byte {
	if l.readPos >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPos]
	}
}
