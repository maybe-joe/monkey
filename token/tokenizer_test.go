package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Tokenizer_Empty(t *testing.T) {
	assert.Equal(t, []Token{Eof()}, NewTokenizer("").Tokenize())
}

func Test_Tokenizer_Next(t *testing.T) {
	tz := NewTokenizer("=+(){},;fn let aAbBcC_")

	testcases := []struct {
		name     string
		expected Token
	}{
		{"Assignment", Assignment()},
		{"Plus", Plus()},
		{"Left Parenthesis", LeftParenthesis()},
		{"Right Parenthesis", RightParenthesis()},
		{"Left Brace", LeftBrace()},
		{"Right Brace", RightBrace()},
		{"Comma", Comma()},
		{"Semicolon", Semicolon()},
		{"Function", Function()},
		{"Let", Let()},
		{"Identifier", Identifier("aAbBcC_")},
		{"Eof", Eof()},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tz.Next())
		})
	}
}
