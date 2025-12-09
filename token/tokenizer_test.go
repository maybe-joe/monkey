package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Tokenizer_Identifier(t *testing.T) {
	assert.Equal(t, "aAbBcC_", NewTokenizer("aAbBcC_ = 0;").Identifier())
}

func Test_Tokenizer_Tokenize_Empty(t *testing.T) {
	expected := []Token{
		Eof(),
	}

	actual := NewTokenizer("").Tokenize()
	assert.Equal(t, expected, actual)
}

func Test_Tokenizer_Tokenize_TokenTypes(t *testing.T) {
	expected := []Token{
		Assignment(),
		Plus(),
		LeftParenthesis(),
		RightParenthesis(),
		LeftBrace(),
		RightBrace(),
		Comma(),
		Semicolon(),
		Eof(),
	}

	actual := NewTokenizer("=+(){},;").Tokenize()
	assert.Equal(t, expected, actual)
}

func Test_Tokenizer_Tokenize_Keywords(t *testing.T) {
	expected := []Token{
		Function(),
		Let(),
		Eof(),
	}

	actual := NewTokenizer("fn let").Tokenize()
	assert.Equal(t, expected, actual)
}

func Test_Tokenizer_Tokenize_Code(t *testing.T) {
	const code = `
		let five = 5;
		let ten = 10;

		let add = fn(x, y) {
		  x + y;
		};

		let result = add(five, ten);
	`
}
