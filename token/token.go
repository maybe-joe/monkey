package token

type TokenType string

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Identifiers + literals
	IDENT TokenType = "IDENT" // add, foobar, x, y, ...
	INT   TokenType = "INT"   // 1343456

	// Operators
	ASSIGN TokenType = "="
	PLUS   TokenType = "+"

	// Delimiters
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"

	LPAREN TokenType = "("
	RPAREN TokenType = ")"
	LBRACE TokenType = "{"
	RBRACE TokenType = "}"

	// Keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
)

type Token struct {
	Type    TokenType
	Literal string
}

func Illegal(literal byte) Token {
	return Token{Type: ILLEGAL, Literal: string(literal)}
}

func Eof() Token {
	return Token{Type: EOF, Literal: ""}
}

func Assignment() Token {
	return Token{Type: ASSIGN, Literal: "="}
}

func Plus() Token {
	return Token{Type: PLUS, Literal: "+"}
}

func LeftParenthesis() Token {
	return Token{Type: LPAREN, Literal: "("}
}

func RightParenthesis() Token {
	return Token{Type: RPAREN, Literal: ")"}
}

func LeftBrace() Token {
	return Token{Type: LBRACE, Literal: "{"}
}

func RightBrace() Token {
	return Token{Type: RBRACE, Literal: "}"}
}

func Comma() Token {
	return Token{Type: COMMA, Literal: ","}
}

func Semicolon() Token {
	return Token{Type: SEMICOLON, Literal: ";"}
}

func Function() Token {
	return Token{Type: FUNCTION, Literal: "FUNCTION"}
}

func Let() Token {
	return Token{Type: LET, Literal: "LET"}
}

func Identifier(literal string) Token {
	return Token{Type: IDENT, Literal: literal}
}

func Integer(literal string) Token {
	return Token{Type: INT, Literal: literal}
}
