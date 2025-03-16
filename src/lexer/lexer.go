package lexer

import (
	"lexicon/src/token"
	"unicode"
)

type Lexer struct {
	input   string // source code
	currPos int    // current charecter position
	nextPos int    // next charecter position
	ch      rune
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
		l.ch = rune(l.input[l.nextPos])
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

func (l *Lexer) readNumber() string {
	start := l.currPos
	hasDot := false

	for unicode.IsDigit(rune(l.ch)) || (l.ch == '.' && !hasDot) {
		if l.ch == '.' {
			// If we already encountered a dot, it's an error
			if hasDot {
				return l.input[start:l.currPos] // Return incomplete number
			}
			hasDot = true
		}
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

// for multi character tokens like !=, ==
func (l *Lexer) peekChar() rune {
	if l.nextPos >= len(l.input) {
		return 0 // EOF
	}
	return rune(l.input[l.nextPos])
}

// creates token.Token to reduce code duplication
func (l *Lexer) newToken(tokenType token.TokenType, ch string) token.Token {
	return token.Token{Type: tokenType, Literal: ch}
}

func (l *Lexer) NextToken() token.Token {

	var tok token.Token
	l.skipWhitespace()

	switch l.ch {
	case '#':
		comment := l.readComment()
		tok = token.Token{Type: token.COMMENT, Literal: comment}
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.EQ, string(ch)+string(l.ch))
		} else {
			tok = l.newToken(token.ASSIGN, "=")
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = l.newToken(token.NOT_EQ, string(ch)+string(l.ch))
		} else {
			tok = l.newToken(token.LOGICAL_NOT, "!")
		}
	case '&':
		if l.peekChar() == '&' {
			l.readChar()
			tok = l.newToken(token.LOGICAL_AND, "&&")
		}
	case '|':
		if l.peekChar() == '|' {
			l.readChar()
			tok = l.newToken(token.LOGICAL_OR, "||")
		}
	case '+':
		tok = l.newToken(token.PLUS, "+")
	case '-':
		tok = l.newToken(token.MINUS, "-")
	case '*':
		tok = l.newToken(token.MUL, "*")
	case '/':
		tok = l.newToken(token.DIV, "/")
	case '%':
		tok = l.newToken(token.MOD, "%")
	case '>':
		tok = l.newToken(token.GT, ">")
	case '<':
		tok = l.newToken(token.LT, "<")
	case '(':
		tok = l.newToken(token.LPAREN, "(")
	case ')':
		tok = l.newToken(token.RPAREN, ")")
	case '{':
		tok = l.newToken(token.LBRACE, "{")
	case '}':
		tok = l.newToken(token.RBRACE, "}")
	case ',':
		tok = l.newToken(token.COMMA, ",")
	case ';':
		tok = l.newToken(token.SEMICOLON, ";")
	case ':':
		tok = l.newToken(token.COLON, ":")
	case '.':
		tok = l.newToken(token.DOT, ".")
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if unicode.IsLetter(rune(l.ch)) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal) // Efficient keyword check
			return tok
		} else if unicode.IsDigit(rune(l.ch)) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = l.newToken(token.ILLEGAL, string(l.ch))
		}
	}

	l.readChar()
	return tok
}
