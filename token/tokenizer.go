package token

import (
	"os"
	"unicode"
)

type Tokenizer struct {
	file   string
	pos    int
	offset int
	curr   byte

	tokens []Token
}

func NewTokenizer(filename string) (*Tokenizer, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Tokenizer{
		file:   string(bytes),
		pos:    0,
		offset: 0,
		curr:   bytes[0],
	}, nil
}

func isAlpha(s string) bool {
  for _, s := range s {
    if (!unicode.IsLetter(s)) {
      return false;
    }
  }

  return true;
}

func isNum(s string) bool {
  for _, s := range s {
    if (!unicode.IsDigit(s)) {
      return false
    }
  }

  return true;
}

func (t *Tokenizer) nextChar() {
  t.pos++
  t.offset++
  t.curr = t.file[t.pos]
}

func (t *Tokenizer) NextToken() *Token {
  for t.curr == ' ' {
    t.nextChar() 
  }

  var nt Token

	switch t.curr {
	case '=':
    nt = Token{ASSIGN, string(t.curr)}
    break
	case ';':
		nt = Token{SEMICOLON, string(t.curr)}
    break
	case '(':
		nt = Token{LPAREN, string(t.curr)}
    break
  case ')':
    nt = Token{RPAREN, string(t.curr)}
    break
	case ',':
    nt = Token{COMMA, string(t.curr)}   		
    break
	case '{':
    nt = Token{LBRACE, string(t.curr)}
    break
	case '}':
    nt = Token{RBRACE, string(t.curr)}
    break
	case '+':
    nt = Token{PLUS, string(t.curr)}
    break
  case 0: 
    nt = Token{EOF, string(t.curr)}
    break
  default:
    if isNum(string(t.curr)) {
      for isNum(string(t.curr)){
        t.offset++ 
      }

      lit := string(t.file[t.pos:t.offset])
      nt = Token{NUMBER, lit}
      t.pos = t.offset
      t.curr = t.file[t.pos]
      return &nt
    }

    if isAlpha(string(t.curr))  {
      for isAlpha(string(t.curr)) || isNum(string(t.curr)) {
        t.offset++
      }

      lit := string(t.file[t.pos:t.offset])
      nt.Literal = lit
      t.pos = t.offset
      t.curr = t.file[t.pos]

      // reserved word check
      switch lit {
      case "let":
        nt.Type = LET
        break
      default:
        nt.Type = IDENT
      }

      return &nt
    }

    nt = Token{ILLEGAL, string(t.curr)}
	}

  t.nextChar()
  return &nt
}

