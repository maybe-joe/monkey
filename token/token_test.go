package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Tokenizer_Tokenize_Empty(t *testing.T) {
	expected := []Token{}

	actual := NewTokenizer("").Tokenize()
	assert.Equal(t, expected, actual)
}

func Test_Tokenizer_Tokenize_kitchensink(t *testing.T) {
	expected := []Token{
		{Type: ASSIGN, Literal: "="},
		{Type: PLUS, Literal: "+"},
		{Type: LPAREN, Literal: "("},
		{Type: RPAREN, Literal: ")"},
		{Type: LBRACE, Literal: "{"},
		{Type: RBRACE, Literal: "}"},
		{Type: COMMA, Literal: ","},
		{Type: SEMICOLON, Literal: ";"},
	}

	actual := NewTokenizer("=+(){},;").Tokenize()
	assert.Equal(t, expected, actual)
}
