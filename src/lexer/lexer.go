package lexer

import (
	"lexicon/src/token"
	"unicode"
)

type Lexer struct {
	input   string // source code
	currPos int    // current charecter position
	nextPos int    // next charecter position
	ch      byte
}

func New(input string) *Lexer {
	l := new(Lexer)
	l.input = input

	l.readChar()
	return l
}

// Reads next character
func (l *Lexer) readChar() {
	if l.nextPos >= len(l.input) {
		l.ch = 0 // EOF
	} else {
		l.ch = l.input[l.nextPos]
	}
	l.currPos = l.nextPos
	l.nextPos++
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(rune(l.ch)) {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	start := l.currPos
	for unicode.IsLetter(rune(l.ch)) {
		l.readChar()
	}
	return l.input[start:l.currPos]
}

func (l *Lexer) readComment() string {
	start := l.currPos
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
	return l.input[start:l.currPos]
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()

	var tok token.Token

	switch l.ch {
	case '#':
		comment := l.readComment()
		tok = token.Token{Type: token.COMMENT, Literal: comment}
	case '{':
		tok = token.Token{Type: token.LBRACE, Literal: "{"}
	case '}':
		tok = token.Token{Type: token.RBRACE, Literal: "}"}
	case '(':
		tok = token.Token{Type: token.LPAREN, Literal: "("}
	case ')':
		tok = token.Token{Type: token.RPAREN, Literal: ")"}
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
	default:
		if unicode.IsLetter(rune(l.ch)) {
			ident := l.readIdentifier()
			if ident == "if" {
				tok = token.Token{Type: token.IF, Literal: ident}
			} else {
				tok = token.Token{Type: token.IDENT, Literal: ident}
			}
			return tok
		} else {
			tok = token.Token{Type: token.ILLEGAL, Literal: string(l.ch)}
		}
	}

	l.readChar()
	return tok
}
