package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"

	// identifiers + literals
	EOF   = "EOF"
	IDENT = "IDENT"
	INT   = "INT"

	// operators
	ASSIGN = "="
	PLUS   = "+"

	// delimeters
	COMMA     = ","
	SEMICOLON = ";"

	LPARENT = "("
	RPARENT = ")"
	LBRACE  = "{"
	RBRACE  = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
