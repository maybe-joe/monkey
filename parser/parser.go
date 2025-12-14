package parser

import (
	"strconv"

	"github.com/maybe-joe/monkey/ast"
	"github.com/maybe-joe/monkey/token"
)

var (
	ErrExpectedIdentifier = "expected identifier after let"
	ErrExpectedAssignment = "expected assignment after identifier"
	ErrExpectedSemicolon  = "expected semicolon after expression"
)

// Order of precedence
const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

// Precedences maps token types to their precedence level
// for infix operators.
var precedences = map[token.TokenType]int{
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.LPAREN:   CALL,
}

type (
	prefixFn func() ast.Expression
	infixFn  func(ast.Expression) ast.Expression
)

type Tokenizer interface {
	Next() token.Token
}

type Parser struct {
	tokenizer Tokenizer
	current   token.Token
	next      token.Token
	errors    []string

	prefixLookup map[token.TokenType]prefixFn
	infixLookup  map[token.TokenType]infixFn
}

func New(tokenizer Tokenizer) *Parser {
	p := &Parser{
		tokenizer: tokenizer,
		errors:    []string{},
	}

	p.prefixLookup = map[token.TokenType]prefixFn{
		token.IDENT:    p.Identifier,
		token.INT:      p.Integer,
		token.BANG:     p.Prefix,
		token.MINUS:    p.Prefix,
		token.TRUE:     p.Boolean,
		token.FALSE:    p.Boolean,
		token.LPAREN:   p.Group,
		token.IF:       p.If,
		token.FUNCTION: p.Function,
	}

	p.infixLookup = map[token.TokenType]infixFn{
		token.PLUS:     p.Infix,
		token.MINUS:    p.Infix,
		token.SLASH:    p.Infix,
		token.ASTERISK: p.Infix,
		token.EQ:       p.Infix,
		token.NOT_EQ:   p.Infix,
		token.LT:       p.Infix,
		token.GT:       p.Infix,
		token.LPAREN:   p.Call,
	}

	p.Next()
	p.Next()

	return p
}

func (p *Parser) Let() *ast.LetNode {
	// If the next token is not an identifier the code is invalid.
	if !p.next.Is(token.IDENT) {
		return nil
	}

	// Advance to the identifier token.
	p.Next()

	// Create the identifier node.
	id := ast.IdentifierNode{
		Value: p.current.Literal,
	}

	// Next we expect an assignment token.
	if !p.next.Is(token.ASSIGN) {
		return nil
	}

	// Advance to the assignment token.
	p.Next()
	p.Next()

	// The next token should be the start of the expression.
	expr := p.Expression(LOWEST)

	// Possibly advance to the semicolon.
	if p.next.Is(token.SEMICOLON) {
		p.Next()
	}

	return &ast.LetNode{
		Identifier: &id,
		Value:      expr,
	}
}

func (p *Parser) Return() *ast.ReturnNode {
	// Advance to the expression token.
	p.Next()

	// The next token should be the start of the expression.
	expr := p.Expression(LOWEST)

	// Possibly advance to the semicolon.
	if p.next.Is(token.SEMICOLON) {
		p.Next()
	}

	return &ast.ReturnNode{
		Value: expr,
	}
}

func (p *Parser) If() ast.Expression {
	if !p.next.Is(token.LPAREN) {
		return nil
	}

	p.Next()
	p.Next()

	condition := p.Expression(LOWEST)

	if !p.next.Is(token.RPAREN) {
		return nil
	}

	p.Next()

	if !p.next.Is(token.LBRACE) {
		return nil
	}

	p.Next()

	consequence := p.Block()

	var alternative *ast.BlockNode
	if p.next.Is(token.ELSE) {
		p.Next()

		if !p.next.Is(token.LBRACE) {
			return nil
		}

		p.Next()

		alternative = p.Block()
	}

	return &ast.IfNode{
		Condition:   condition,
		Consequence: consequence,
		Alternative: alternative,
	}
}

func (p *Parser) Block() *ast.BlockNode {
	stmts := []ast.Statement{}

	p.Next()

	for !p.current.Is(token.RBRACE) && !p.current.Is(token.EOF) {
		if stmt := p.Statement(); stmt != nil {
			stmts = append(stmts, stmt)
		}
		p.Next()
	}

	return &ast.BlockNode{
		Statements: stmts,
	}
}

func (p *Parser) Identifier() ast.Expression {
	return &ast.IdentifierNode{
		Value: p.current.Literal,
	}
}

func (p *Parser) Integer() ast.Expression {
	i, err := strconv.ParseInt(p.current.Literal, 0, 64)
	if err != nil {
		p.errors = append(p.errors, "could not parse "+p.current.Literal+" as integer")
		return nil
	}

	return &ast.IntegerNode{
		Value: i,
	}
}

func (p *Parser) Boolean() ast.Expression {
	return &ast.BooleanNode{
		Value: p.current.Is(token.TRUE),
	}
}

func (p *Parser) Call(function ast.Expression) ast.Expression {
	return &ast.CallNode{
		Function:  function,
		Arguments: p.Arguments(),
	}
}

func (p *Parser) Arguments() []ast.Expression {
	if p.next.Is(token.RPAREN) {
		p.Next()
		return nil
	}

	p.Next()

	args := []ast.Expression{
		p.Expression(LOWEST),
	}

	for p.next.Is(token.COMMA) {
		p.Next()
		p.Next()
		args = append(args, p.Expression(LOWEST))
	}

	if !p.next.Is(token.RPAREN) {
		return nil
	}

	p.Next()
	return args
}

func (p *Parser) Function() ast.Expression {
	if !p.next.Is(token.LPAREN) {
		return nil
	}

	p.Next()

	parameters := p.Parameters()

	if !p.next.Is(token.LBRACE) {
		return nil
	}

	p.Next()

	body := p.Block()

	return &ast.FunctionNode{
		Parameters: parameters,
		Body:       body,
	}
}

func (p *Parser) Parameters() []*ast.IdentifierNode {
	if p.next.Is(token.RPAREN) {
		p.Next()
		return nil
	}

	p.Next()

	identifiers := []*ast.IdentifierNode{
		{Value: p.current.Literal},
	}

	for p.next.Is(token.COMMA) {
		p.Next()
		p.Next()
		identifiers = append(identifiers, &ast.IdentifierNode{Value: p.current.Literal})
	}

	if !p.next.Is(token.RPAREN) {
		return nil
	}

	p.Next()

	return identifiers
}

func (p *Parser) Group() ast.Expression {
	p.Next()

	expr := p.Expression(LOWEST)

	if !p.next.Is(token.RPAREN) {
		return nil
	}

	p.Next()

	return expr
}

func (p *Parser) Prefix() ast.Expression {
	expr := &ast.PrefixNode{
		Operator: p.current.String(),
	}

	p.Next()

	expr.Right = p.Expression(PREFIX)

	return expr
}

func (p *Parser) Infix(left ast.Expression) ast.Expression {
	expr := &ast.InfixNode{
		Left:     left,
		Operator: p.current.String(),
	}

	precedence := LOWEST
	if val, ok := precedences[p.current.Type]; ok {
		precedence = val
	}

	p.Next()
	expr.Right = p.Expression(precedence)

	return expr
}

func (p *Parser) Expression(precedence int) ast.Expression {
	prefix, ok := p.prefixLookup[p.current.Type]
	if !ok {
		return nil
	}

	expr := prefix()

	for !p.next.Is(token.SEMICOLON) && precedence < precedences[p.next.Type] {
		infix, ok := p.infixLookup[p.next.Type]
		if !ok {
			return expr
		}

		p.Next()

		expr = infix(expr)
	}

	return expr
}

func (p *Parser) ExpressionStatement() *ast.ExpressionStatementNode {
	expr := p.Expression(LOWEST)

	if p.next.Is(token.SEMICOLON) {
		p.Next()
	}

	return &ast.ExpressionStatementNode{
		Expression: expr,
	}
}

func (p *Parser) Statement() ast.Statement {
	switch p.current.Type {
	default:
		return p.ExpressionStatement()
	case token.LET:
		return p.Let()
	case token.RETURN:
		return p.Return()
	}
}

func (p *Parser) Next() {
	p.current = p.next
	p.next = p.tokenizer.Next()
}

func (p *Parser) Until(typ token.TokenType) {
	for !p.current.Is(typ) && !p.current.Is(token.EOF) {
		p.Next()
	}
}

func (p *Parser) Parse() *ast.RootNode {
	root := &ast.RootNode{}

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
