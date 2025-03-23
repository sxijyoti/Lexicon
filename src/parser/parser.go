package parser

import (
	"lexicon/src/ast"
	"lexicon/src/lexer"
	"lexicon/src/token"
)

type Parser struct {
	l         *lexer.Lexer
	currToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekToken.Type == t {
		p.nextToken()
		return true
	}
	return false
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.IDENT:
		return p.parseAssignment()
	case token.SPROUT:
		return p.parseVariableDeclaration()
	case token.IF:
		return p.parseIfExpression()
	default:
		return nil
	}
}

// sprout x = 10 | sprout x int = 10
func (p *Parser) parseVariableDeclaration() *ast.VariableDeclaration {
	stmt := &ast.VariableDeclaration{Token: p.currToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}

	if p.peekToken.Type == token.IDENT {
		p.nextToken()
		stmt.Type = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
	}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()
	stmt.Value = &ast.IntegerLiteral{Token: p.currToken} // TODO: parse actual expressions
	return stmt
}

// x = 10
func (p *Parser) parseAssignment() *ast.VariableDeclaration {
	stmt := &ast.VariableDeclaration{Token: p.currToken}
	stmt.Name = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()
	stmt.Value = &ast.IntegerLiteral{Token: p.currToken} // TODO: parse actual expressions
	return stmt
}

// if (cond) { } else { }
func (p *Parser) parseIfExpression() *ast.IfExpression {
	expr := &ast.IfExpression{Token: p.currToken}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}
	p.nextToken()
	expr.Condition = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	expr.Consequence = p.parseBlockStatement()

	if p.peekToken.Type == token.ELSE {
		p.nextToken()
		if !p.expectPeek(token.LBRACE) {
			return nil
		}
		expr.Alternative = p.parseBlockStatement()
	}

	return expr
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.currToken}
	p.nextToken()

	for p.currToken.Type != token.RBRACE && p.currToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}

	return block
}
