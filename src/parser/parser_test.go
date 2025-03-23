package parser

import (
	"lexicon/src/ast"
	"lexicon/src/lexer"
	"testing"
)

func TestVariableDeclaration(t *testing.T) {
	input := `
	sprout x = 10;
	sprout y int = 20;
	x = 30;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if len(program.Statements) != 3 {
		t.Fatalf("expected 3 statements, got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.VariableDeclaration)
	if !ok {
		t.Fatalf("expected *ast.VariableDeclaration, got=%T", program.Statements[0])
	}
	if stmt.Name.Value != "x" {
		t.Errorf("expected variable name 'x', got=%s", stmt.Name.Value)
	}
}

func TestIfElseParsing(t *testing.T) {
	input := `
	if (x) {
		sprout y = 10;
	} else {
		sprout z = 20;
	}
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if len(program.Statements) != 1 {
		t.Fatalf("expected 1 statement, got=%d", len(program.Statements))
	}

	ifStmt, ok := program.Statements[0].(*ast.IfExpression)
	if !ok {
		t.Fatalf("expected *ast.IfExpression, got=%T", program.Statements[0])
	}

	ident, ok := ifStmt.Condition.(*ast.Identifier)
	if !ok || ident.Value != "x" {
		t.Errorf("expected condition to be identifier 'x', got=%v", ifStmt.Condition)
	}

	if len(ifStmt.Consequence.Statements) != 1 {
		t.Errorf("expected 1 statement in consequence, got=%d", len(ifStmt.Consequence.Statements))
	}

	if len(ifStmt.Alternative.Statements) != 1 {
		t.Errorf("expected 1 statement in alternative, got=%d", len(ifStmt.Alternative.Statements))
	}
}
