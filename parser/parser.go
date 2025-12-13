package parser

import (
	"github.com/maybe-joe/monkey/token"
)

var (
	ErrExpectedIdentifier = "expected identifier after let"
	ErrExpectedAssignment = "expected assignment after identifier"
	ErrExpectedSemicolon  = "expected semicolon after expression"
)

type Tokenizer interface {
	Next() token.Token
}

type Parser struct {
	tokenizer Tokenizer
	current   token.Token
	next      token.Token
	errors    []string
}

func New(tokenizer Tokenizer) *Parser {
	p := &Parser{
		tokenizer: tokenizer,
		errors:    []string{},
	}

	p.Next()
	p.Next()

	return p
}

func (p *Parser) Let() *Let {
	// If the next token is not an identifier the code is invalid.
	if !p.next.Is(token.IDENT) {
		p.errors = append(p.errors, ErrExpectedIdentifier)
		return nil
	}

	// Advance to the identifier token.
	p.Next()

	// Create the identifier node.
	id := Identifier{
		Value: p.current.Literal,
	}

	// Next we expect an assignment token.
	if !p.next.Is(token.ASSIGN) {
		p.errors = append(p.errors, ErrExpectedAssignment)
		return nil
	}

	// Advance to the assignment token.
	p.Next()

	// For now we will skip the expression until we reach a semicolon.
	for !p.next.Is(token.SEMICOLON) || p.next.Is(token.EOF) {
		p.Next()
	}

	// Advance to the semicolon.
	p.Next()

	if !p.current.Is(token.SEMICOLON) {
		p.errors = append(p.errors, ErrExpectedSemicolon)
		return nil
	}

	return &Let{
		Identifier: &id,
	}
}

func (p *Parser) Statement() Statement {
	switch p.current.Type {
	default:
		return nil
	case token.LET:
		return p.Let()
	}
}

func (p *Parser) Next() {
	p.current = p.next
	p.next = p.tokenizer.Next()
}

func (p *Parser) Parse() *Root {
	root := &Root{}

	for p.current.Type != token.EOF {
		if stmt := p.Statement(); stmt != nil {
			root.Statements = append(root.Statements, stmt)
		}

		p.Next()
	}

	return root
}

func (p *Parser) Errors() []string {
	return p.errors
}
