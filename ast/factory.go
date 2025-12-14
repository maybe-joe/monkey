package ast

func Root(statements ...Statement) *RootNode {
	return &RootNode{Statements: statements}
}

func Let(identifier *IdentifierNode, value Expression) *LetNode {
	return &LetNode{Identifier: identifier, Value: value}
}

func Return(value Expression) *ReturnNode {
	return &ReturnNode{Value: value}
}

func If(condition Expression, consequence *BlockNode, alternative *BlockNode) *IfNode {
	return &IfNode{Condition: condition, Consequence: consequence, Alternative: alternative}
}

func Block(statements ...Statement) *BlockNode {
	return &BlockNode{Statements: statements}
}

func Function(parameters []*IdentifierNode, body *BlockNode) *FunctionNode {
	return &FunctionNode{Parameters: parameters, Body: body}
}

func Identifier(value string) *IdentifierNode {
	return &IdentifierNode{Value: value}
}

// TODO: Do the parsing here
func Integer(value int64) *IntegerNode {
	return &IntegerNode{Value: value}
}

func True() *BooleanNode {
	return &BooleanNode{Value: true}
}

func False() *BooleanNode {
	return &BooleanNode{Value: false}
}

func Call(function Expression, arguments ...Expression) *CallNode {
	return &CallNode{Function: function, Arguments: arguments}
}

func ExpressionStatement(expression Expression) *ExpressionStatementNode {
	return &ExpressionStatementNode{Expression: expression}
}

func Prefix(operator string, right Expression) *PrefixNode {
	return &PrefixNode{Operator: operator, Right: right}
}

func Infix(left Expression, operator string, right Expression) *InfixNode {
	return &InfixNode{Left: left, Operator: operator, Right: right}
}
