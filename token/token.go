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

func (t Token) String() string {
	if len(t.Literal) > 0 {
		return string(t.Type) + " " + t.Literal
	} else {
		return string(t.Type)
	}
}

func Illegal(literal byte) Token {
	return Token{Type: ILLEGAL, Literal: string(literal)}
}

func Eof() Token {
	return Token{Type: EOF}
}

func Assignment() Token {
	return Token{Type: ASSIGN}
}

func Plus() Token {
	return Token{Type: PLUS}
}

func Minus() Token {
	return Token{Type: MINUS}
}

func Bang() Token {
	return Token{Type: BANG}
}

func Asterisk() Token {
	return Token{Type: ASTERISK}
}

func Slash() Token {
	return Token{Type: SLASH}
}

func LessThan() Token {
	return Token{Type: LT}
}

func GreaterThan() Token {
	return Token{Type: GT}
}

func Equal() Token {
	return Token{Type: EQ}
}

func NotEqual() Token {
	return Token{Type: NOT_EQ}
}

func LeftParenthesis() Token {
	return Token{Type: LPAREN}
}

func RightParenthesis() Token {
	return Token{Type: RPAREN}
}

func LeftBrace() Token {
	return Token{Type: LBRACE}
}

func RightBrace() Token {
	return Token{Type: RBRACE}
}

func Comma() Token {
	return Token{Type: COMMA}
}

func Semicolon() Token {
	return Token{Type: SEMICOLON}
}

func Function() Token {
	return Token{Type: FUNCTION}
}

func Let() Token {
	return Token{Type: LET}
}

func True() Token {
	return Token{Type: TRUE}
}

func False() Token {
	return Token{Type: FALSE}
}

func If() Token {
	return Token{Type: IF}
}

func Else() Token {
	return Token{Type: ELSE}
}

func Return() Token {
	return Token{Type: RETURN}
}

func Integer(literal string) Token {
	return Token{Type: INT, Literal: literal}
}

func Identifier(literal string) Token {
	return Token{Type: IDENT, Literal: literal}
}
