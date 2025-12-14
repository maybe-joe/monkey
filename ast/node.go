package ast

type Node interface {
	node()
}

type Statement interface {
	Node
	statement()
}

type Expression interface {
	Node
	expression()
}

type RootNode struct {
	Statements []Statement
}

func (RootNode) node()      {}
func (RootNode) statement() {}

type LetNode struct {
	Identifier *IdentifierNode
	Value      Expression
}

func (LetNode) node()      {}
func (LetNode) statement() {}

type ReturnNode struct {
	Value Expression
}

func (ReturnNode) node()      {}
func (ReturnNode) statement() {}

type IfNode struct {
	Condition   Expression
	Consequence *BlockNode
	Alternative *BlockNode
}

func (IfNode) node()       {}
func (IfNode) expression() {}

type BlockNode struct {
	Statements []Statement
}

func (BlockNode) node()      {}
func (BlockNode) statement() {}

type FunctionNode struct {
	Parameters []*IdentifierNode
	Body       *BlockNode
}

func (FunctionNode) node()       {}
func (FunctionNode) expression() {}

type IdentifierNode struct {
	Value string
}

func (IdentifierNode) node()       {}
func (IdentifierNode) expression() {}

type IntegerNode struct {
	Value int64
}

func (IntegerNode) node()       {}
func (IntegerNode) expression() {}

type BooleanNode struct {
	Value bool
}

func (BooleanNode) node()       {}
func (BooleanNode) expression() {}

type CallNode struct {
	Function  Expression
	Arguments []Expression
}

func (CallNode) node()       {}
func (CallNode) expression() {}

type ExpressionStatementNode struct {
	Expression Expression
}

func (ExpressionStatementNode) node()      {}
func (ExpressionStatementNode) statement() {}

type PrefixNode struct {
	Operator string
	Right    Expression
}

func (PrefixNode) node()       {}
func (PrefixNode) expression() {}

type InfixNode struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (InfixNode) node()       {}
func (InfixNode) expression() {}
