package lexer

import (
	"fmt"
	"strings"
	"testing"

	"github.com/SimonRichardson/runscript/internal/token"
)

func TestNextToken(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []token.Token
	}{
		{
			name:  "empty",
			input: "",
			expected: []token.Token{{
				Type:  token.EOF,
				Start: token.Position{Line: 1, Column: 0},
				End:   token.Position{Line: 1, Column: 0},
			}},
		},
		{
			name:  "single",
			input: "foo",
			expected: []token.Token{{
				Type:    token.IDENT,
				Literal: "foo",
				Start:   token.Position{Line: 1, Column: 0},
				End:     token.Position{Line: 1, Column: 3},
			}},
		},
		{
			name:  "multiple",
			input: "foo bar",
			expected: []token.Token{{
				Type:    token.IDENT,
				Literal: "foo",
				Start:   token.Position{Line: 1, Column: 0},
				End:     token.Position{Line: 1, Column: 3},
			}, {
				Type:    token.IDENT,
				Literal: "bar",
				Start:   token.Position{Line: 1, Column: 4},
				End:     token.Position{Line: 1, Column: 7},
			}},
		},
		{
			name: "newline",
			input: `foo
bar`,
			expected: []token.Token{{
				Type:    token.IDENT,
				Literal: "foo",
				Start:   token.Position{Line: 1, Column: 0},
				End:     token.Position{Line: 1, Column: 3},
			}, {
				Type:    token.IDENT,
				Literal: "bar",
				Start:   token.Position{Line: 2, Column: 0},
				End:     token.Position{Line: 2, Column: 3},
			}},
		},
		{
			name: "newline+space",
			input: `foo		    bar
baz       qux`,
			expected: []token.Token{{
				Type:    token.IDENT,
				Literal: "foo",
				Start:   token.Position{Line: 1, Column: 0},
				End:     token.Position{Line: 1, Column: 3},
			}, {
				Type:    token.IDENT,
				Literal: "bar",
				Start:   token.Position{Line: 1, Column: 9},
				End:     token.Position{Line: 1, Column: 12},
			}, {
				Type:    token.IDENT,
				Literal: "baz",
				Start:   token.Position{Line: 2, Column: 0},
				End:     token.Position{Line: 2, Column: 3},
			}, {
				Type:    token.IDENT,
				Literal: "qux",
				Start:   token.Position{Line: 2, Column: 10},
				End:     token.Position{Line: 2, Column: 13},
			}},
		},
		{
			name:  "characters",
			input: "a b 1 = 2.1 ! , ; ( ) & | [ ]",
			expected: []token.Token{{
				Type:    token.IDENT,
				Literal: "a",
				Start:   token.Position{Line: 1, Column: 0},
				End:     token.Position{Line: 1, Column: 1},
			}, {
				Type:    token.IDENT,
				Literal: "b",
				Start:   token.Position{Line: 1, Column: 2},
				End:     token.Position{Line: 1, Column: 3},
			}, {
				Type:    token.INT,
				Literal: "1",
				Start:   token.Position{Line: 1, Column: 4},
				End:     token.Position{Line: 1, Column: 5},
			}, {
				Type:    token.ASSIGN,
				Literal: "=",
				Start:   token.Position{Line: 1, Column: 6},
				End:     token.Position{Line: 1, Column: 7},
			}, {
				Type:    token.FLOAT,
				Literal: "2.1",
				Start:   token.Position{Line: 1, Column: 8},
				End:     token.Position{Line: 1, Column: 11},
			}, {
				Type:    token.BANG,
				Literal: "!",
				Start:   token.Position{Line: 1, Column: 12},
				End:     token.Position{Line: 1, Column: 13},
			}, {
				Type:    token.COMMA,
				Literal: ",",
				Start:   token.Position{Line: 1, Column: 14},
				End:     token.Position{Line: 1, Column: 15},
			}, {
				Type:    token.SEMICOLON,
				Literal: ";",
				Start:   token.Position{Line: 1, Column: 16},
				End:     token.Position{Line: 1, Column: 17},
			}, {
				Type:    token.LPAREN,
				Literal: "(",
				Start:   token.Position{Line: 1, Column: 18},
				End:     token.Position{Line: 1, Column: 19},
			}, {
				Type:    token.RPAREN,
				Literal: ")",
				Start:   token.Position{Line: 1, Column: 20},
				End:     token.Position{Line: 1, Column: 21},
			}, {
				Type:    token.BITAND,
				Literal: "&",
				Start:   token.Position{Line: 1, Column: 22},
				End:     token.Position{Line: 1, Column: 23},
			}, {
				Type:    token.BITOR,
				Literal: "|",
				Start:   token.Position{Line: 1, Column: 24},
				End:     token.Position{Line: 1, Column: 25},
			}, {
				Type:    token.LBRACKET,
				Literal: "[",
				Start:   token.Position{Line: 1, Column: 26},
				End:     token.Position{Line: 1, Column: 27},
			}, {
				Type:    token.RBRACKET,
				Literal: "]",
				Start:   token.Position{Line: 1, Column: 28},
				End:     token.Position{Line: 1, Column: 29},
			}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			l := New(strings.NewReader(c.input))

			for _, e := range c.expected {
				tok := l.Next()

				t.Run(fmt.Sprintf("literal_%s", e.Literal), func(t *testing.T) {

					if tok.Type != e.Type {
						t.Errorf("expected type %d, got %d", e.Type, tok.Type)
					} else if tok.Literal != e.Literal {
						t.Errorf("expected literal %q, got %q", e.Literal, tok.Literal)
					} else if tok.Start.Line != e.Start.Line {
						t.Errorf("expected start line %d, got %d", e.Start.Line, tok.Start.Line)
					} else if tok.Start.Column != e.Start.Column {
						t.Errorf("expected start column %d, got %d", e.Start.Column, tok.Start.Column)
					} else if tok.End.Line != e.End.Line {
						t.Errorf("expected end line %d, got %d", e.End.Line, tok.End.Line)
					} else if tok.End.Column != e.End.Column {
						t.Errorf("expected end column %d, got %d", e.End.Column, tok.End.Column)
					}
				})
			}
		})
	}
}

func TestNextTokenNumbers(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []token.Token
	}{
		{
			name:  "zero",
			input: "0",
			expected: []token.Token{{
				Type:    token.INT,
				Start:   token.Position{Line: 1, Column: 0},
				End:     token.Position{Line: 1, Column: 1},
				Literal: "0",
			}},
		},
		{
			name:  "zero",
			input: "0.0",
			expected: []token.Token{{
				Type:    token.FLOAT,
				Start:   token.Position{Line: 1, Column: 0},
				End:     token.Position{Line: 1, Column: 3},
				Literal: "0.0",
			}},
		},
		{
			name:  "zero",
			input: ".0",
			expected: []token.Token{{
				Type:    token.FLOAT,
				Start:   token.Position{Line: 1, Column: 0},
				End:     token.Position{Line: 1, Column: 2},
				Literal: ".0",
			}},
		},
		{
			name:  "big",
			input: "4222222222.9999999999999999999",
			expected: []token.Token{{
				Type:    token.FLOAT,
				Start:   token.Position{Line: 1, Column: 0},
				End:     token.Position{Line: 1, Column: 30},
				Literal: "4222222222.9999999999999999999",
			}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			l := New(strings.NewReader(c.input))

			for _, e := range c.expected {
				tok := l.Next()

				t.Run(fmt.Sprintf("literal_%s", e.Literal), func(t *testing.T) {

					if tok.Type != e.Type {
						t.Errorf("expected type %d, got %d", e.Type, tok.Type)
					} else if tok.Literal != e.Literal {
						t.Errorf("expected literal %q, got %q", e.Literal, tok.Literal)
					} else if tok.Start.Line != e.Start.Line {
						t.Errorf("expected start line %d, got %d", e.Start.Line, tok.Start.Line)
					} else if tok.Start.Column != e.Start.Column {
						t.Errorf("expected start column %d, got %d", e.Start.Column, tok.Start.Column)
					} else if tok.End.Line != e.End.Line {
						t.Errorf("expected end line %d, got %d", e.End.Line, tok.End.Line)
					} else if tok.End.Column != e.End.Column {
						t.Errorf("expected end column %d, got %d", e.End.Column, tok.End.Column)
					}
				})
			}
		})
	}
}
