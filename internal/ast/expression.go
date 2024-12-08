package ast

import (
	"bytes"
	"strings"

	"github.com/SimonRichardson/runscript/internal/token"
)

// Expression defines a type of AST node for outlining an expression.
type Expression interface {
	Positions() []token.Position

	String() string
}

// QueryExpression represents a query full of expressions
type QueryExpression struct {
	Expressions []Expression
}

func (e QueryExpression) Positions() []token.Position {
	if len(e.Expressions) > 0 {
		var positions []token.Position
		for _, expr := range e.Expressions {
			positions = append(positions, expr.Positions()...)
		}
		return positions
	}
	return nil
}

func (e QueryExpression) String() string {
	var out bytes.Buffer

	for _, s := range e.Expressions {
		out.WriteString(s.String())
	}

	return out.String()
}

// ExpressionStatement is a group of expressions that allows us to group a
// subset of expressions.
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es ExpressionStatement) Positions() []token.Position {
	return []token.Position{es.Token.Start, es.Token.End}
}

func (es ExpressionStatement) String() string {
	if es.Expression != nil {
		str := es.Expression.String()
		if str == "" {
			return ";"
		}
		if str[len(str)-1:] != ";" {
			str += ";"
		}
		return strings.TrimSpace(str)
	}
	return ""
}

// Identifier represents an identifier for a given AST block
type Identifier struct {
	Token token.Token
}

// Positions returns the positions of the identifier.
func (i Identifier) Positions() []token.Position {
	return []token.Position{i.Token.Start, i.Token.End}
}

func (i Identifier) String() string { return i.Token.Literal }
