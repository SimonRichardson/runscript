package parser

import (
	"reflect"
	"strings"
	"testing"

	"github.com/SimonRichardson/runscript/internal/ast"
	"github.com/SimonRichardson/runscript/internal/lexer"
	"github.com/SimonRichardson/runscript/internal/token"
)

func TestParser(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected ast.QueryExpression
	}{
		{
			name:  "empty",
			input: "",
			expected: ast.QueryExpression{
				Expressions: []ast.Expression{
					ast.ExpressionStatement{
						Token: token.Token{
							Type:  token.EOF,
							Start: token.Position{Line: 1, Column: 0},
							End:   token.Position{Line: 1, Column: 0},
						},
					},
				},
			},
		},
		{
			name:  "single ident",
			input: "foo",
			expected: ast.QueryExpression{
				Expressions: []ast.Expression{
					ast.Identifier{
						Token: token.Token{
							Type:    token.IDENT,
							Literal: "foo",
							Start:   token.Position{Line: 1, Column: 0},
							End:     token.Position{Line: 1, Column: 3},
						},
					},
					ast.ExpressionStatement{
						Token: token.Token{
							Type:  token.EOF,
							Start: token.Position{Line: 1, Column: 0},
							End:   token.Position{Line: 1, Column: 0},
						},
					},
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			lexer := lexer.New(strings.NewReader(tc.input))

			p := New(lexer)
			expr, err := p.Parse()
			if err != nil {
				t.Fatal(err)
			}
			if reflect.DeepEqual(expr, tc.expected) {
				t.Fatalf("expected %v, got %v", tc.expected, expr)
			}
		})
	}
}
