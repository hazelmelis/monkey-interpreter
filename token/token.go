package token

type TokenType string

// Token data structure has fields for its type and literal value
type Token struct {
	Type    TokenType
	Literal string
}

// Define possible TokenTypes in Monkey as constants
const (
	ILLEGAL = "ILLEGAL" // Unknown token/character
	EOF     = "EOF"     // "End of file", tells parser to stop

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
