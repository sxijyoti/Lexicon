package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Datatypes for Variable Declaration
	IDENT      = "IDENT"
	INT        = "INT"
	STRING     = "STRING"
	FLOAT      = "FLOAT"
	TYPE_IDENT = "TYPE_IDENT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	MUL    = "*"
	DIV    = "/"
	MOD    = "%"

	// Comparison Operators
	GT     = ">"
	LT     = "<"
	EQ     = "=="
	NOT_EQ = "!="

	// Bitwise Operators
	// XOR = "^"
	// AND = "&"
	// OR  = "|"
	// NOT = "!"

	// Logical Operators
	LOGICAL_AND = "AND"
	LOGICAL_OR  = "OR"
	LOGICAL_NOT = "NOT"

	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	COLON     = ":"
	DOT       = "."

	// Keywords
	ECHO   = "ECHO"
	SPROUT = "SPROUT"
)

var keywords = map[string]TokenType{
	"echo":   ECHO,
	"sprout": SPROUT,

	"and": LOGICAL_AND,
	"or":  LOGICAL_OR,
	"not": LOGICAL_NOT,

	"int":    TYPE_IDENT,
	"float":  TYPE_IDENT,
	"string": TYPE_IDENT,
}

// To check if the identifier is a keyword or not
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
