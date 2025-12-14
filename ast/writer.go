package ast

import (
	"fmt"
	"io"
	"strings"
)

type Writer struct {
	writer io.Writer
	indent int
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{writer: w, indent: 0}
}

func (w *Writer) Write(node Node) {
	switch n := node.(type) {
	case *IntegerNode:
		w.Integer(n)
	case *BooleanNode:
		w.Boolean(n)
	case *IdentifierNode:
		w.Identifier(n)
	case *BlockNode:
		w.Block(n)
	case *ReturnNode:
		w.Return(n)
	case *LetNode:
		w.Let(n)
	case *CallNode:
		w.Call(n)
	case *PrefixNode:
		w.Prefix(n)
	case *InfixNode:
		w.Infix(n)
	case *FunctionNode:
		w.Function(n)
	case *IfNode:
		w.If(n)
	case *ExpressionStatementNode:
		w.ExpressionStatement(n)
	case *RootNode:
		for _, stmt := range n.Statements {
			w.Write(stmt)
		}
	default:
		fmt.Fprintf(w.writer, "<%T>", n)
	}
}

func (w *Writer) Indentation() string {
	return strings.Repeat("\t", w.indent)
}

func (w *Writer) Integer(node *IntegerNode) {
	fmt.Fprintf(w.writer, "%d", node.Value)
}

func (w *Writer) Boolean(node *BooleanNode) {
	if node.Value {
		fmt.Fprint(w.writer, "true")
	} else {
		fmt.Fprint(w.writer, "false")
	}
}

func (w *Writer) Identifier(node *IdentifierNode) {
	fmt.Fprint(w.writer, node.Value)
}

func (w *Writer) Block(node *BlockNode) {
	fmt.Fprint(w.writer, "{\n")
	w.indent++
	for _, stmt := range node.Statements {
		fmt.Fprint(w.writer, w.Indentation())
		w.Write(stmt)
	}
	w.indent--
	fmt.Fprint(w.writer, "\n}")
}

func (w *Writer) Return(node *ReturnNode) {
	fmt.Fprint(w.writer, "return ")
	w.Write(node.Value)
	fmt.Fprint(w.writer, ";")
}

func (w *Writer) Let(node *LetNode) {
	fmt.Fprint(w.writer, "let ")
	w.Identifier(node.Identifier)
	fmt.Fprint(w.writer, " = ")
	w.Write(node.Value)
	fmt.Fprint(w.writer, ";\n")
}

func (w *Writer) Call(node *CallNode) {
	w.Write(node.Function)
	fmt.Fprint(w.writer, "(")
	for i, arg := range node.Arguments {
		w.Write(arg)
		if i < len(node.Arguments)-1 {
			fmt.Fprint(w.writer, ", ")
		}
	}
	fmt.Fprint(w.writer, ")")
}

func (w *Writer) Prefix(node *PrefixNode) {
	fmt.Fprintf(w.writer, "%s", node.Operator)
	w.Write(node.Right)
}

func (w *Writer) Infix(node *InfixNode) {
	fmt.Fprint(w.writer, "(")
	w.Write(node.Left)
	fmt.Fprintf(w.writer, " %s ", node.Operator)
	w.Write(node.Right)
	fmt.Fprint(w.writer, ")")
}

func (w *Writer) Function(node *FunctionNode) {
	fmt.Fprint(w.writer, "fn(")
	for i, param := range node.Parameters {
		w.Identifier(param)
		if i < len(node.Parameters)-1 {
			fmt.Fprint(w.writer, ", ")
		}
	}
	fmt.Fprint(w.writer, ") ")
	w.Block(node.Body)
}

func (w *Writer) If(node *IfNode) {
	fmt.Fprint(w.writer, "if ")
	w.Write(node.Condition)
	fmt.Fprint(w.writer, " ")
	w.Block(node.Consequence)
	if node.Alternative != nil {
		fmt.Fprint(w.writer, " else ")
		w.Block(node.Alternative)
	}
}

func (w *Writer) ExpressionStatement(node *ExpressionStatementNode) {
	w.Write(node.Expression)
}
