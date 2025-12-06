package token

type TokenType string

type Token struct {
	Type	TokenType
	Literal	string
}

const {
	ILLEGAL = "ILLEGAL"  // token not defined below
	EOF 	= "EOF"		 // parser can stop
	
	IDENT = "IDENT"		 // variables, function names
	INT	  = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA	  = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET		 = "LET"
}
