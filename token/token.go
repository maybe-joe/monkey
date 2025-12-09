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

func Illegal() Token {
	return Token{Type: ILLEGAL, Literal: ""}
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

// Tokenizer, converts a string of monkey code into tokens.
type Tokenizer struct {
	// code, program string to tokenize.
	code string
	// char, current char under examination.
	// If the end of code is reached, 0 is used to represent 'EOF'.
	char byte
	// cursor, the index of char in the code.
	cursor int
	// peek, one char lookahead.
	peek int
}

// NewTokenizer creates a new Tokenizer for the given code.
func NewTokenizer(code string) *Tokenizer {
	tz := &Tokenizer{
		code:   code,
		cursor: 0,
		peek:   1,
		char:   0, // EOF
	}

	// Prevent reading out of bounds.
	if len(tz.code) > 0 {
		tz.char = tz.code[0]
	}

	return tz
}

// Advance the tokenizer to the next character.
func (tz *Tokenizer) Advance() {
	if tz.peek >= len(tz.code) {
		tz.char = 0
	} else {
		tz.char = tz.code[tz.peek]
	}

	tz.cursor = tz.peek
	tz.peek++
}

// Next returns the next token from the code and advances the tokenizer.
func (tz *Tokenizer) Next() Token {
	var t Token

	switch tz.char {
	default:
		t = Illegal()
	case 0:
		t = Eof()
	case '=':
		t = Assignment()
	case '+':
		t = Plus()
	case '(':
		t = LeftParenthesis()
	case ')':
		t = RightParenthesis()
	case '{':
		t = LeftBrace()
	case '}':
		t = RightBrace()
	case ',':
		t = Comma()
	case ';':
		t = Semicolon()
	}

	tz.Advance()
	return t
}

// Tokenize processes the entire input code and returns a slice of tokens.
func (tz *Tokenizer) Tokenize() []Token {
	tokens := make([]Token, 0, 100)

	for {
		t := tz.Next()
		if t.Type == EOF {
			break
		}

		tokens = append(tokens, t)
	}

	return tokens
}
