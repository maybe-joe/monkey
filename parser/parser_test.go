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

	expected := &ast.Root{
		Statements: []ast.Statement{
			&ast.Let{
				Identifier: &ast.Identifier{Value: "x"},
				Value:      &ast.Integer{Value: 5},
			},
			&ast.Let{
				Identifier: &ast.Identifier{Value: "y"},
				Value:      &ast.Integer{Value: 10},
			},
			&ast.Let{
				Identifier: &ast.Identifier{Value: "foobar"},
				Value:      &ast.Integer{Value: 838383},
			},
			&ast.Return{
				Value: &ast.Integer{Value: 5},
			},
			&ast.ExpressionStatement{
				Expression: &ast.Identifier{Value: "foobar"},
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

	expected := &ast.Root{
		Statements: []ast.Statement{
			&ast.Let{
				Identifier: &ast.Identifier{Value: "x"},
				Value:      &ast.Integer{Value: 67},
			},
			&ast.Let{
				Identifier: &ast.Identifier{Value: "y"},
				Value:      &ast.Identifier{Value: "x"},
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

	expected := &ast.Root{
		Statements: []ast.Statement{
			&ast.Return{
				Value: &ast.Integer{Value: 67},
			},
			&ast.Return{
				Value: &ast.Identifier{Value: "x"},
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

	expected := &ast.Root{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.Prefix{
					Operator: "!",
					Right:    &ast.Integer{Value: 5},
				},
			},
			&ast.ExpressionStatement{
				Expression: &ast.Prefix{
					Operator: "-",
					Right:    &ast.Integer{Value: 15},
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

	expected := &ast.Root{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.Infix{
					Left:     &ast.Integer{Value: 5},
					Operator: "+",
					Right:    &ast.Integer{Value: 5},
				},
			},
			&ast.ExpressionStatement{
				Expression: &ast.Infix{
					Left:     &ast.Integer{Value: 10},
					Operator: "-",
					Right:    &ast.Integer{Value: 2},
				},
			},
			&ast.ExpressionStatement{
				Expression: &ast.Infix{
					Left:     &ast.Integer{Value: 3},
					Operator: "*",
					Right:    &ast.Integer{Value: 4},
				},
			},
			&ast.ExpressionStatement{
				Expression: &ast.Infix{
					Left:     &ast.Integer{Value: 8},
					Operator: "/",
					Right:    &ast.Integer{Value: 2},
				},
			},
			&ast.ExpressionStatement{
				Expression: &ast.Infix{
					Left:     &ast.Integer{Value: 5},
					Operator: ">",
					Right:    &ast.Integer{Value: 3},
				},
			},
			&ast.ExpressionStatement{
				Expression: &ast.Infix{
					Left:     &ast.Integer{Value: 2},
					Operator: "<",
					Right:    &ast.Integer{Value: 4},
				},
			},
			&ast.ExpressionStatement{
				Expression: &ast.Infix{
					Left:     &ast.Integer{Value: 6},
					Operator: "==",
					Right:    &ast.Integer{Value: 6},
				},
			},
			&ast.ExpressionStatement{
				Expression: &ast.Infix{
					Left:     &ast.Integer{Value: 7},
					Operator: "!=",
					Right:    &ast.Integer{Value: 9},
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

	expected := &ast.Root{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.Boolean{Value: true},
			},
			&ast.ExpressionStatement{
				Expression: &ast.Boolean{Value: false},
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

	expected := &ast.Root{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.Infix{
					Left: &ast.Infix{
						Left:     &ast.Integer{Value: 5},
						Operator: "+",
						Right:    &ast.Integer{Value: 5},
					},
					Operator: "*",
					Right:    &ast.Integer{Value: 2},
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

	expected := &ast.Root{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.If{
					Condition: &ast.Infix{
						Left:     &ast.Identifier{Value: "x"},
						Operator: "<",
						Right:    &ast.Identifier{Value: "y"},
					},
					Consequence: &ast.Block{
						Statements: []ast.Statement{
							&ast.ExpressionStatement{
								Expression: &ast.Identifier{Value: "x"},
							},
						},
					},
					Alternative: &ast.Block{
						Statements: []ast.Statement{
							&ast.ExpressionStatement{
								Expression: &ast.Identifier{Value: "y"},
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

	expected := &ast.Root{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.Function{
					Parameters: []*ast.Identifier{
						{Value: "x"},
						{Value: "y"},
					},
					Body: &ast.Block{
						Statements: []ast.Statement{
							&ast.ExpressionStatement{
								Expression: &ast.Infix{
									Left:     &ast.Identifier{Value: "x"},
									Operator: "+",
									Right:    &ast.Identifier{Value: "y"},
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

	expected := &ast.Root{
		Statements: []ast.Statement{
			&ast.ExpressionStatement{
				Expression: &ast.Call{
					Function: &ast.Identifier{Value: "add"},
					Arguments: []ast.Expression{
						&ast.Integer{Value: 1},
						&ast.Infix{
							Left:     &ast.Integer{Value: 2},
							Operator: "+",
							Right:    &ast.Integer{Value: 3},
						},
						&ast.Infix{
							Left:     &ast.Integer{Value: 4},
							Operator: "*",
							Right:    &ast.Integer{Value: 5},
						},
					},
				},
			},
		},
	}

	actual := New(token.NewTokenizer(given)).Parse()
	assert.Equal(t, expected, actual)
}
