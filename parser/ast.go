package parser

type Node interface {
	Literal() string
}

type Statement interface {
	statement()
	Literal() string
}

type Expression interface {
	expression()
	Literal() string
}

type Root struct {
	Statements []Statement
}

func (n *Root) Literal() string {
	if len(n.Statements) > 0 {
		return n.Statements[0].Literal()
	} else {
		return ""
	}
}

type Let struct {
	Identifier *Identifier
	Value      Expression
}

func (n *Let) statement() {}

func (n *Let) Literal() string {
	return "let"
}

type Identifier struct {
	Value string
}

func (n *Identifier) expression() {}

func (n *Identifier) Literal() string {
	return n.Value
}
