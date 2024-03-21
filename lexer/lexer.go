package lexer

import (
	"os"
	"unicode"

	"github.com/slinky55/milo/token"
)

type Lexer struct {
	file   string
	pos    int
	offset int
	curr   byte

	tokens []token.Token
}

func New(filename string) (*Lexer, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Lexer{
		file:   string(bytes),
		pos:    0,
		offset: 0,
		curr:   bytes[0],
	}, nil
}

func isAlpha(b byte) bool {
	return unicode.IsLetter(rune(b))
}

func isNum(b byte) bool {
	return unicode.IsDigit(rune(b))
}

func isWhitespace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\n' || b == '\r'
}

func (l *Lexer) nextChar() {
	l.pos++
	l.offset++
	if l.pos >= len(l.file) {
		return
	}
	l.curr = l.file[l.pos]
}

func (l *Lexer) NextToken() *token.Token {
	if l.pos >= len(l.file) {
		return &token.Token{
			Type:    token.EOF,
			Literal: "",
		}
	}

	for isWhitespace(l.curr) {
		l.nextChar()
		if l.pos >= len(l.file) {
			return &token.Token{
				Type:    token.EOF,
				Literal: "",
			}
		}
	}

	var nt token.Token

	switch l.curr {
	case '=':
		nt = token.Token{
			Type:    token.ASSIGN,
			Literal: string(l.curr),
		}
		break
	case ';':
		nt = token.Token{
			Type:    token.SEMICOLON,
			Literal: string(l.curr),
		}
		break
	case '(':
		nt = token.Token{
			Type:    token.LPAREN,
			Literal: string(l.curr),
		}
		break
	case ')':
		nt = token.Token{
			Type:    token.RPAREN,
			Literal: string(l.curr),
		}
		break
	case ',':
		nt = token.Token{
			Type:    token.COMMA,
			Literal: string(l.curr),
		}
		break
	case '{':
		nt = token.Token{
			Type:    token.LBRACE,
			Literal: string(l.curr),
		}
		break
	case '}':
		nt = token.Token{
			Type:    token.RBRACE,
			Literal: string(l.curr),
		}
		break
	case '+':
		nt = token.Token{
			Type:    token.PLUS,
			Literal: string(l.curr),
		}
		break
	case 0:
		return &token.Token{
			Type:    token.EOF,
			Literal: string(l.curr),
		}
	default:
		if isNum(l.curr) {
			for isNum(l.file[l.offset]) {
				l.offset++
			}

			lit := l.file[l.pos:l.offset]
			nt = token.Token{
				Type:    token.NUMBER,
				Literal: lit,
			}
			l.pos = l.offset
			l.curr = l.file[l.pos]
			return &nt
		}

		if isAlpha(l.curr) {
			for isAlpha(l.file[l.offset]) || isNum(l.file[l.offset]) {
				l.offset++
			}

			lit := l.file[l.pos:l.offset]
			nt.Literal = lit
			l.pos = l.offset
			l.curr = l.file[l.pos]

			// reserved word check
			switch lit {
			case "let":
				nt.Type = token.LET
				break
			case "fn":
				nt.Type = token.FUNCTION
				break
			case "const":
				nt.Type = token.CONST
				break
			case "int":
				nt.Type = token.INT
				break
			case "return":
				nt.Type = token.RETURN
				break
			default:
				nt.Type = token.IDENT
			}

			return &nt
		}

		nt = token.Token{
			Type:    token.ILLEGAL,
			Literal: string(l.curr),
		}
	}

	l.nextChar()
	return &nt
}
