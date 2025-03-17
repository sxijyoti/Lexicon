package lexer

import (
	"lexicon/src/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `sprout x = 5; 
			if (x > 2) { 
			x = x + 1; 
			echo 1
			}`

	expectedTokens := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.SPROUT, "sprout"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.GT, ">"},
		{token.INT, "2"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.INT, "1"},
		{token.SEMICOLON, ";"},
		{token.ECHO, "echo"},
		{token.INT, "1"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, expected := range expectedTokens {
		tok := lexer.NextToken()

		if tok.Type != expected.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q",
				i, expected.expectedType, tok.Type)
		}

		if tok.Literal != expected.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, expected.expectedLiteral, tok.Literal)
		}
	}
}
