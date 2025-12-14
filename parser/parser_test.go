package parser

import (
	"testing"

	"github.com/maybe-joe/monkey/ast"
	"github.com/maybe-joe/monkey/token"
	"github.com/stretchr/testify/assert"
)

func Test_Program(t *testing.T) {
	given := `
		let x = 5;
		let y = 10;
		let foobar = 838383;

		return 5;
		foobar;
	`

	expected := &ast.RootNode{
		Statements: []ast.Statement{
			&ast.LetNode{
				Identifier: &ast.IdentifierNode{Value: "x"},
				Value:      &ast.IntegerNode{Value: 5},
			},
			&ast.LetNode{
				Identifier: &ast.IdentifierNode{Value: "y"},
				Value:      &ast.IntegerNode{Value: 10},
			},
			&ast.LetNode{
				Identifier: &ast.IdentifierNode{Value: "foobar"},
				Value:      &ast.IntegerNode{Value: 838383},
			},
			&ast.ReturnNode{
				Value: &ast.IntegerNode{Value: 5},
			},
			&ast.ExpressionStatementNode{
				Expression: &ast.IdentifierNode{Value: "foobar"},
			},
		},
	}

	actual := New(token.NewTokenizer(given)).Parse()
	assert.Equal(t, expected, actual)
}

func Test_Let(t *testing.T) {
	given := `
		let x = 67;
		let y = x;
	`

	expected := &ast.RootNode{
		Statements: []ast.Statement{
			&ast.LetNode{
				Identifier: &ast.IdentifierNode{Value: "x"},
				Value:      &ast.IntegerNode{Value: 67},
			},
			&ast.LetNode{
				Identifier: &ast.IdentifierNode{Value: "y"},
				Value:      &ast.IdentifierNode{Value: "x"},
			},
		},
	}

	actual := New(token.NewTokenizer(given)).Parse()
	assert.Equal(t, expected, actual)
}

func Test_Return(t *testing.T) {
	given := `
		return 67;
		return x;
	`

	expected := &ast.RootNode{
		Statements: []ast.Statement{
			&ast.ReturnNode{
				Value: &ast.IntegerNode{Value: 67},
			},
			&ast.ReturnNode{
				Value: &ast.IdentifierNode{Value: "x"},
			},
		},
	}

	actual := New(token.NewTokenizer(given)).Parse()
	assert.Equal(t, expected, actual)
}

func Test_Prefix(t *testing.T) {
	given := `
		!5;
		-15;
	`

	expected := &ast.RootNode{
		Statements: []ast.Statement{
			&ast.ExpressionStatementNode{
				Expression: &ast.PrefixNode{
					Operator: "!",
					Right:    &ast.IntegerNode{Value: 5},
				},
			},
			&ast.ExpressionStatementNode{
				Expression: &ast.PrefixNode{
					Operator: "-",
					Right:    &ast.IntegerNode{Value: 15},
				},
			},
		},
	}

	actual := New(token.NewTokenizer(given)).Parse()
	assert.Equal(t, expected, actual)
}

func Test_Infix(t *testing.T) {
	given := `
		5 + 5;
		10 - 2;
		3 * 4;
		8 / 2;
		5 > 3;
		2 < 4;
		6 == 6;
		7 != 9;
	`

	expected := &ast.RootNode{
		Statements: []ast.Statement{
			&ast.ExpressionStatementNode{
				Expression: &ast.InfixNode{
					Left:     &ast.IntegerNode{Value: 5},
					Operator: "+",
					Right:    &ast.IntegerNode{Value: 5},
				},
			},
			&ast.ExpressionStatementNode{
				Expression: &ast.InfixNode{
					Left:     &ast.IntegerNode{Value: 10},
					Operator: "-",
					Right:    &ast.IntegerNode{Value: 2},
				},
			},
			&ast.ExpressionStatementNode{
				Expression: &ast.InfixNode{
					Left:     &ast.IntegerNode{Value: 3},
					Operator: "*",
					Right:    &ast.IntegerNode{Value: 4},
				},
			},
			&ast.ExpressionStatementNode{
				Expression: &ast.InfixNode{
					Left:     &ast.IntegerNode{Value: 8},
					Operator: "/",
					Right:    &ast.IntegerNode{Value: 2},
				},
			},
			&ast.ExpressionStatementNode{
				Expression: &ast.InfixNode{
					Left:     &ast.IntegerNode{Value: 5},
					Operator: ">",
					Right:    &ast.IntegerNode{Value: 3},
				},
			},
			&ast.ExpressionStatementNode{
				Expression: &ast.InfixNode{
					Left:     &ast.IntegerNode{Value: 2},
					Operator: "<",
					Right:    &ast.IntegerNode{Value: 4},
				},
			},
			&ast.ExpressionStatementNode{
				Expression: &ast.InfixNode{
					Left:     &ast.IntegerNode{Value: 6},
					Operator: "==",
					Right:    &ast.IntegerNode{Value: 6},
				},
			},
			&ast.ExpressionStatementNode{
				Expression: &ast.InfixNode{
					Left:     &ast.IntegerNode{Value: 7},
					Operator: "!=",
					Right:    &ast.IntegerNode{Value: 9},
				},
			},
		},
	}

	actual := New(token.NewTokenizer(given)).Parse()
	assert.Equal(t, expected, actual)
}

func Test_Boolean(t *testing.T) {
	given := `
		true;
		false;
	`

	expected := &ast.RootNode{
		Statements: []ast.Statement{
			&ast.ExpressionStatementNode{
				Expression: &ast.BooleanNode{Value: true},
			},
			&ast.ExpressionStatementNode{
				Expression: &ast.BooleanNode{Value: false},
			},
		},
	}

	actual := New(token.NewTokenizer(given)).Parse()
	assert.Equal(t, expected, actual)
}

func Test_Group(t *testing.T) {
	given := `
		(5 + 5) * 2;
	`

	expected := &ast.RootNode{
		Statements: []ast.Statement{
			&ast.ExpressionStatementNode{
				Expression: &ast.InfixNode{
					Left: &ast.InfixNode{
						Left:     &ast.IntegerNode{Value: 5},
						Operator: "+",
						Right:    &ast.IntegerNode{Value: 5},
					},
					Operator: "*",
					Right:    &ast.IntegerNode{Value: 2},
				},
			},
		},
	}

	actual := New(token.NewTokenizer(given)).Parse()
	assert.Equal(t, expected, actual)
}

func Test_If(t *testing.T) {
	given := `
		if (x < y) { x } else { y }
	`

	expected := &ast.RootNode{
		Statements: []ast.Statement{
			&ast.ExpressionStatementNode{
				Expression: &ast.IfNode{
					Condition: &ast.InfixNode{
						Left:     &ast.IdentifierNode{Value: "x"},
						Operator: "<",
						Right:    &ast.IdentifierNode{Value: "y"},
					},
					Consequence: &ast.BlockNode{
						Statements: []ast.Statement{
							&ast.ExpressionStatementNode{
								Expression: &ast.IdentifierNode{Value: "x"},
							},
						},
					},
					Alternative: &ast.BlockNode{
						Statements: []ast.Statement{
							&ast.ExpressionStatementNode{
								Expression: &ast.IdentifierNode{Value: "y"},
							},
						},
					},
				},
			},
		},
	}

	actual := New(token.NewTokenizer(given)).Parse()
	assert.Equal(t, expected, actual)
}

func Test_Function(t *testing.T) {
	given := `
		fn(x, y) { x + y; }
	`

	expected := &ast.RootNode{
		Statements: []ast.Statement{
			&ast.ExpressionStatementNode{
				Expression: &ast.FunctionNode{
					Parameters: []*ast.IdentifierNode{
						{Value: "x"},
						{Value: "y"},
					},
					Body: &ast.BlockNode{
						Statements: []ast.Statement{
							&ast.ExpressionStatementNode{
								Expression: &ast.InfixNode{
									Left:     &ast.IdentifierNode{Value: "x"},
									Operator: "+",
									Right:    &ast.IdentifierNode{Value: "y"},
								},
							},
						},
					},
				},
			},
		},
	}

	actual := New(token.NewTokenizer(given)).Parse()
	assert.Equal(t, expected, actual)
}

func Test_Call(t *testing.T) {
	given := `
		add(1, 2 + 3, 4 * 5);
	`

	expected := &ast.RootNode{
		Statements: []ast.Statement{
			&ast.ExpressionStatementNode{
				Expression: &ast.CallNode{
					Function: &ast.IdentifierNode{Value: "add"},
					Arguments: []ast.Expression{
						&ast.IntegerNode{Value: 1},
						&ast.InfixNode{
							Left:     &ast.IntegerNode{Value: 2},
							Operator: "+",
							Right:    &ast.IntegerNode{Value: 3},
						},
						&ast.InfixNode{
							Left:     &ast.IntegerNode{Value: 4},
							Operator: "*",
							Right:    &ast.IntegerNode{Value: 5},
						},
					},
				},
			},
		},
	}

	actual := New(token.NewTokenizer(given)).Parse()
	assert.Equal(t, expected, actual)
}
