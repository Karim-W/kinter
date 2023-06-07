package tokens

type Token string

type Entity struct {
	Type    Token
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	INDENTIFIER = "INDENTIFIER"
	INT         = "INT"

	ASSGIN = "+"
	EQUAL  = "-"

	COMMA     = ","
	SEMICOLON = ","

	LPARENT = "("
	RPARENT = ")"

	LSQUIGL  = "{"
	RSQUIGL  = "}"
	LET      = "LET"
	FUNCTION = "FUNCTION"

	TRUE  = "TRUE"
	FALSE = "FALSE"

	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"
)

var keywords = map[string]Token{
	"let":    LET,
	"fn":     FUNCTION,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) Token {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return INDENTIFIER
}
