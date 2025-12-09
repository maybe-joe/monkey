package token

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

// Peek, returns the next character without advancing the tokenizer.
func (tz *Tokenizer) Peek() byte {
	if tz.peek >= len(tz.code) {
		return 0
	} else {
		return tz.code[tz.peek]
	}
}

// Whitespace, skips over whitespace characters.
func (tz *Tokenizer) Whitespace() {
	for isWhitespace(tz.char) {
		tz.Advance()
	}
}

// Identifier, reads an identifier from the code and returns it as a Token.
func (tz *Tokenizer) Identifier() string {
	start := tz.cursor

	for isLetter(tz.char) {
		tz.Advance()
	}

	return tz.code[start:tz.cursor]
}

// Number, reads a number from the code and returns it as a Token.
func (tz *Tokenizer) Number() string {
	start := tz.cursor

	for isDigit(tz.char) {
		tz.Advance()
	}

	return tz.code[start:tz.cursor]
}

// Next returns the next token from the code and advances the tokenizer.
func (tz *Tokenizer) Next() Token {
	var t Token

	// Skip whitespace
	tz.Whitespace()

	switch tz.char {
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
	default:
		if isLetter(tz.char) {
			switch literal := tz.Identifier(); literal {
			default:
				return Identifier(literal)
			case "fn":
				return Function()
			case "let":
				return Let()
			}
		} else if isDigit(tz.char) {
			return Integer(tz.Number())
		} else {
			t = Illegal(tz.char)
		}
	}

	tz.Advance()
	return t
}

// Tokenize processes the entire input code and returns a slice of tokens.
func (tz *Tokenizer) Tokenize() []Token {
	tokens := make([]Token, 0, 100)

	for {
		t := tz.Next()
		tokens = append(tokens, t)

		// Stop if we reach EOF, but include it in the tokens.
		if t.Type == EOF {
			break
		}
	}

	return tokens
}

// isWhitespace, returns true if the given character is a whitespace character.
func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

// isLetter, returns true if the given character is a letter or underscore.
// These are valid characters for identifiers in Monkey.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit, returns true if the given character is a digit.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
