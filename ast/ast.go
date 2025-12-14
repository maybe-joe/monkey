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

type Root struct {
	Statements []Statement
}

func (Root) node()      {}
func (Root) statement() {}

type Let struct {
	Identifier *Identifier
	Value      Expression
}

func (Let) node()      {}
func (Let) statement() {}

type Return struct {
	Value Expression
}

func (Return) node()      {}
func (Return) statement() {}

type If struct {
	Condition   Expression
	Consequence *Block
	Alternative *Block
}

func (If) node()       {}
func (If) expression() {}

type Block struct {
	Statements []Statement
}

func (Block) node()      {}
func (Block) statement() {}

type Function struct {
	Parameters []*Identifier
	Body       *Block
}

func (Function) node()       {}
func (Function) expression() {}

type Identifier struct {
	Value string
}

func (Identifier) node()       {}
func (Identifier) expression() {}

type Integer struct {
	Value int64
}

func (Integer) node()       {}
func (Integer) expression() {}

type Boolean struct {
	Value bool
}

func (Boolean) node()       {}
func (Boolean) expression() {}

type Call struct {
	Function  Expression
	Arguments []Expression
}

func (Call) node()       {}
func (Call) expression() {}

type ExpressionStatement struct {
	Expression Expression
}

func (ExpressionStatement) node()      {}
func (ExpressionStatement) statement() {}

type Prefix struct {
	Operator string
	Right    Expression
}

func (Prefix) node()       {}
func (Prefix) expression() {}

type Infix struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (Infix) node()       {}
func (Infix) expression() {}
