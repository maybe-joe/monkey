package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Tokenizer_Empty(t *testing.T) {
	assert.Equal(t, []Token{Eof()}, NewTokenizer("").Tokenize())
}

func Test_Tokenizer_Next(t *testing.T) {
	tz := NewTokenizer("=+(){},;fn let aAbBcC_ 9 1!-/*<>")

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
		{"Bang", Bang()},
		{"Minus", Minus()},
		{"Slash", Slash()},
		{"Asterisk", Asterisk()},
		{"Less Than", LessThan()},
		{"Greater Than", GreaterThan()},
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

		if (5 < 10) {
			return true;
		} else {
			return false;
		}
	`

	expected := []Token{
		Let(), Identifier("five"), Assignment(), Integer("5"), Semicolon(),
		Let(), Identifier("ten"), Assignment(), Integer("10"), Semicolon(),
		Let(), Identifier("add"), Assignment(), Function(), LeftParenthesis(),
		Identifier("x"), Comma(), Identifier("y"), RightParenthesis(),
		LeftBrace(),
		Identifier("x"), Plus(), Identifier("y"), Semicolon(),
		RightBrace(), Semicolon(),
		Let(), Identifier("result"), Assignment(),
		Identifier("add"), LeftParenthesis(),
		Identifier("five"), Comma(), Identifier("ten"), RightParenthesis(),
		Semicolon(),
		If(), LeftParenthesis(), Integer("5"), LessThan(), Integer("10"), RightParenthesis(),
		LeftBrace(),
		Return(), True(), Semicolon(),
		RightBrace(),
		Else(),
		LeftBrace(),
		Return(), False(), Semicolon(),
		RightBrace(),
		Eof(),
	}

	assert.Equal(t, expected, NewTokenizer(code).Tokenize())
}
