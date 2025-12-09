package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Tokenizer_Empty(t *testing.T) {
	assert.Equal(t, []Token{Eof()}, NewTokenizer("").Tokenize())
}

func Test_Tokenizer_Next(t *testing.T) {
	tz := NewTokenizer("=+(){},;fn let aAbBcC_ 9 1")

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
		{"9", Integer("9")},
		{"1", Integer("1")},
		{"Eof", Eof()},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tz.Next())
		})
	}
}

func Test_Tokenizer_Tokenize_SimpleProgram(t *testing.T) {
	const code = `
		let five = 5;
		let ten = 10;

		let add = fn(x, y) {
		  x + y;
		};

		let result = add(five, ten);
	`
}
