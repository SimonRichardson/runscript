package parser

import (
	"fmt"

	"github.com/SimonRichardson/runscript/internal/token"
)

type SyntaxError struct {
	Positions []token.Position
	TokenType token.TokenType
}

func (e SyntaxError) Error() string {
	return fmt.Sprintf("syntax error: unexpected %s", e.TokenType)
}
