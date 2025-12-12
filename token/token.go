package token

type TokenType string

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Identifiers + literals
	IDENT TokenType = "IDENT" // add, foobar, x, y, ...
	INT   TokenType = "INT"   // 1343456

	// Operators
	ASSIGN   TokenType = "="
	PLUS     TokenType = "+"
	MINUS    TokenType = "-"
	BANG     TokenType = "!"
	ASTERISK TokenType = "*"
	SLASH    TokenType = "/"

	// Comparisons
	LT     TokenType = "<"
	GT     TokenType = ">"
	EQ     TokenType = "=="
	NOT_EQ TokenType = "!="

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
	TRUE     TokenType = "TRUE"
	FALSE    TokenType = "FALSE"
	IF       TokenType = "IF"
	ELSE     TokenType = "ELSE"
	RETURN   TokenType = "RETURN"
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

func Minus() Token {
	return Token{Type: MINUS, Literal: "-"}
}

func Bang() Token {
	return Token{Type: BANG, Literal: "!"}
}

func Asterisk() Token {
	return Token{Type: ASTERISK, Literal: "*"}
}

func Slash() Token {
	return Token{Type: SLASH, Literal: "/"}
}

func LessThan() Token {
	return Token{Type: LT, Literal: "<"}
}

func GreaterThan() Token {
	return Token{Type: GT, Literal: ">"}
}

func Equal() Token {
	return Token{Type: EQ, Literal: "=="}
}

func NotEqual() Token {
	return Token{Type: NOT_EQ, Literal: "!="}
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

func Integer(literal string) Token {
	return Token{Type: INT, Literal: literal}
}

func Identifier(literal string) Token {
	return Token{Type: IDENT, Literal: literal}
}

func Function() Token {
	return Token{Type: FUNCTION, Literal: "FUNCTION"}
}

func Let() Token {
	return Token{Type: LET, Literal: "LET"}
}

func True() Token {
	return Token{Type: TRUE, Literal: "TRUE"}
}

func False() Token {
	return Token{Type: FALSE, Literal: "FALSE"}
}

func If() Token {
	return Token{Type: IF, Literal: "IF"}
}

func Else() Token {
	return Token{Type: ELSE, Literal: "ELSE"}
}

func Return() Token {
	return Token{Type: RETURN, Literal: "RETURN"}
}
