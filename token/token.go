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

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// check table to see if string is a keyword or user-defined identifier
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
