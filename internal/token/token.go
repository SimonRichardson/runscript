package token

// TokenType represents a way to identify an individual token.
type TokenType int

const (
	UNKNOWN TokenType = (iota - 1)
	EOF

	IDENT
	INT   //int literal
	FLOAT //float literal
	STRING

	EQ     // ==
	NEQ    // !=
	ASSIGN // =
	BANG   // !

	LT // <
	LE // <=
	GT // >
	GE // >=

	COMMA     // ,
	COLON     // :
	SEMICOLON // ;

	LPAREN   // (
	RPAREN   // )
	LBRACKET // [
	RBRACKET // ]
	LBRACE   // {
	RBRACE   // }

	BITAND  // &
	BITOR   // |
	CONDAND // &&
	CONDOR  // ||
	BOOL    // BOOL

	PLUS     // +
	MINUS    // -
	ASTERISK // *
	SLASH    // /
	MOD      // %

	LAMBDA     // =>
	UNDERSCORE // _
	PERIOD     // .
)

func (t TokenType) String() string {
	switch t {
	case EOF:
		return "EOF"
	case IDENT:
		return "IDENT"
	case INT:
		return "INT"
	case FLOAT:
		return "FLOAT"
	case ASSIGN:
		return "="
	case BANG:
		return "!"
	case EQ:
		return "=="
	case NEQ:
		return "!="
	case LT:
		return "<"
	case LE:
		return "<="
	case GT:
		return ">"
	case GE:
		return ">="
	case LAMBDA:
		return "=>"
	case UNDERSCORE:
		return "_"
	case PERIOD:
		return "."
	case COMMA:
		return ","
	case SEMICOLON:
		return ";"
	case LPAREN:
		return "("
	case RPAREN:
		return ")"
	case LBRACKET:
		return "["
	case RBRACKET:
		return "]"
	case BITAND:
		return "&"
	case BITOR:
		return "|"
	case CONDAND:
		return "&&"
	case CONDOR:
		return "||"
	case BOOL:
		return "BOOL"
	case PLUS:
		return "+"
	case MINUS:
		return "-"
	case ASTERISK:
		return "*"
	case SLASH:
		return "/"
	case MOD:
		return "%"
	case LBRACE:
		return "{"
	case RBRACE:
		return "}"
	default:
		return "<UNKNOWN>"
	}
}

type Position struct {
	Line   int
	Column int
}

type Token struct {
	Type    TokenType
	Start   Position
	End     Position
	Literal string
}
