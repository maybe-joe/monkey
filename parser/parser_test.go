package parser

import (
	"testing"

	"github.com/maybe-joe/monkey/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Program(t *testing.T) {
	given := `
		let x = 5;
		let y = 10;
		let foobar = 838383;

		return 5;
	`

	expected := &Root{
		Statements: []Statement{
			&Let{
				Identifier: &Identifier{Value: "x"},
				// Value:      &IntegerLiteral{Value: 5},
			},
			&Let{
				Identifier: &Identifier{Value: "y"},
				// Value:      &IntegerLiteral{Value: 10},
			},
			&Let{
				Identifier: &Identifier{Value: "foobar"},
				// Value:      &IntegerLiteral{Value: 838383},
			},
			&Return{
				// Value: &IntegerLiteral{Value: 5},
			},
		},
	}

	actual := New(token.NewTokenizer(given)).Parse()
	assert.Equal(t, expected, actual)
}

func Test_Let(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		given := `
			let x = 5;
		`

		expected := &Root{
			Statements: []Statement{
				&Let{
					Identifier: &Identifier{Value: "x"},
					// Value:      &IntegerLiteral{Value: 5},
				},
			},
		}

		actual := New(token.NewTokenizer(given)).Parse()
		assert.Equal(t, expected, actual)
	})

	t.Run("expected identifier", func(t *testing.T) {
		given := `
			let = 5;
		`

		p := New(token.NewTokenizer(given))
		_ = p.Parse()

		require.Len(t, p.errors, 1)
		assert.Equal(t, ErrExpectedIdentifier, p.errors[0])
	})

	t.Run("expected assignment", func(t *testing.T) {
		given := `
			let x 5;
		`

		p := New(token.NewTokenizer(given))
		_ = p.Parse()

		require.Len(t, p.errors, 1)
		assert.Equal(t, ErrExpectedAssignment, p.errors[0])
	})

	t.Run("expected semicolon", func(t *testing.T) {
		given := `
			let x = 5
		`

		p := New(token.NewTokenizer(given))
		_ = p.Parse()

		require.Len(t, p.errors, 1)
		assert.Equal(t, ErrExpectedSemicolon, p.errors[0])
	})
}

func Test_Return(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		given := `
			return 5;
		`

		expected := &Root{
			Statements: []Statement{
				&Return{
					// Value: &IntegerLiteral{Value: 5},
				},
			},
		}

		actual := New(token.NewTokenizer(given)).Parse()
		assert.Equal(t, expected, actual)
	})
}
