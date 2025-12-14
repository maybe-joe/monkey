package ast

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Writer_Write(t *testing.T) {
	testcases := []struct {
		name     string
		given    Node
		expected string
	}{
		{name: "integer", given: Integer(5), expected: "5"},
		{name: "true", given: True(), expected: "true"},
		{name: "false", given: False(), expected: "false"},
		{name: "identifier", given: Identifier("foobar"), expected: "foobar"},
		{name: "block", given: Block(Return(Identifier("x"))), expected: "{return x;}"},
		{name: "return", given: Return(Identifier("x")), expected: "return x;"},
		{name: "let", given: Let(Identifier("x"), Integer(10)), expected: "let x = 10;"},
		{name: "call", given: Call(Identifier("add"), Integer(1), Integer(2)), expected: "add(1, 2)"},
		{name: "prefix", given: Prefix("-", Integer(5)), expected: "-5"},
		{name: "infix", given: Infix(Integer(5), "+", Integer(5)), expected: "(5 + 5)"},
		{name: "function", given: Function(Block(Return(Identifier("x"))), Identifier("x")), expected: "fn(x) {return x;}"},
		{name: "if", given: If(Infix(Identifier("x"), "<", Integer(10)), Block(Return(True())), Block(Return(False()))), expected: "if (x < 10) {return true;} else {return false;}"},
		{name: "expression statement", given: ExpressionStatement(Infix(Integer(5), "+", Integer(5))), expected: "(5 + 5)"},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			NewWriter(&buf).Write(tc.given)
			assert.Equal(t, tc.expected, buf.String())
		})
	}
}

func Test_Writer_Simple(t *testing.T) {
	given := Root(
		Let(
			Identifier("add"),
			Function(
				Block(
					Return(
						Infix(
							Identifier("x"),
							"+",
							Identifier("y"),
						),
					),
				),
				Identifier("x"),
				Identifier("y"),
			),
		),
		Let(
			Identifier("invert"),
			Function(
				Block(
					Return(
						Prefix(
							"-",
							Identifier("x"),
						),
					),
				),
				Identifier("x"),
			),
		),
		Let(
			Identifier("result"),
			Call(
				Identifier("add"),
				Call(Identifier("invert"), Integer(5)),
				Call(Identifier("invert"), Integer(10)),
			),
		),
	)

	var buf bytes.Buffer
	NewWriter(&buf).Write(given)

	expected := `let add = fn(x, y) {return (x + y);};let invert = fn(x) {return -x;};let result = add(invert(5), invert(10));`
	assert.Equal(t, expected, buf.String())
}
