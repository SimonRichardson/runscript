package lexer

import (
	"fmt"
	"strings"
	"testing"
)

func TestNextToken(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []Token
	}{
		{
			name:     "empty",
			input:    "",
			expected: []Token{},
		},
		{
			name:  "single",
			input: "foo",
			expected: []Token{{
				Literal: "foo",
				Start:   Position{Line: 1, Column: 0},
				End:     Position{Line: 1, Column: 3},
			}},
		},
		{
			name:  "multiple",
			input: "foo bar",
			expected: []Token{{
				Literal: "foo",
				Start:   Position{Line: 1, Column: 0},
				End:     Position{Line: 1, Column: 3},
			}, {
				Literal: "bar",
				Start:   Position{Line: 1, Column: 4},
				End:     Position{Line: 1, Column: 7},
			}},
		},
		{
			name: "newline",
			input: `foo
bar`,
			expected: []Token{{
				Literal: "foo",
				Start:   Position{Line: 1, Column: 0},
				End:     Position{Line: 1, Column: 3},
			}, {
				Literal: "bar",
				Start:   Position{Line: 2, Column: 0},
				End:     Position{Line: 2, Column: 3},
			}},
		},
		{
			name: "newline+space",
			input: `foo		    bar
baz       qux`,
			expected: []Token{{
				Literal: "foo",
				Start:   Position{Line: 1, Column: 0},
				End:     Position{Line: 1, Column: 3},
			}, {
				Literal: "bar",
				Start:   Position{Line: 1, Column: 9},
				End:     Position{Line: 1, Column: 12},
			}, {
				Literal: "baz",
				Start:   Position{Line: 2, Column: 0},
				End:     Position{Line: 2, Column: 3},
			}, {
				Literal: "qux",
				Start:   Position{Line: 2, Column: 10},
				End:     Position{Line: 2, Column: 13},
			}},
		},
		{
			name:  "characters",
			input: "a b 1 = 2.1 ! , ; ( ) & | [ ]",
			expected: []Token{{
				Literal: "a",
				Start:   Position{Line: 1, Column: 0},
				End:     Position{Line: 1, Column: 1},
			}, {
				Literal: "b",
				Start:   Position{Line: 1, Column: 2},
				End:     Position{Line: 1, Column: 3},
			}, {
				Literal: "1",
				Start:   Position{Line: 1, Column: 4},
				End:     Position{Line: 1, Column: 5},
			}, {
				Literal: "=",
				Start:   Position{Line: 1, Column: 6},
				End:     Position{Line: 1, Column: 7},
			}, {
				Literal: "2.1",
				Start:   Position{Line: 1, Column: 8},
				End:     Position{Line: 1, Column: 11},
			}, {
				Literal: "!",
				Start:   Position{Line: 1, Column: 12},
				End:     Position{Line: 1, Column: 13},
			}, {
				Literal: ",",
				Start:   Position{Line: 1, Column: 14},
				End:     Position{Line: 1, Column: 15},
			}, {
				Literal: ";",
				Start:   Position{Line: 1, Column: 16},
				End:     Position{Line: 1, Column: 17},
			}, {
				Literal: "(",
				Start:   Position{Line: 1, Column: 18},
				End:     Position{Line: 1, Column: 19},
			}, {
				Literal: ")",
				Start:   Position{Line: 1, Column: 20},
				End:     Position{Line: 1, Column: 21},
			}, {
				Literal: "&",
				Start:   Position{Line: 1, Column: 22},
				End:     Position{Line: 1, Column: 23},
			}, {
				Literal: "|",
				Start:   Position{Line: 1, Column: 24},
				End:     Position{Line: 1, Column: 25},
			}, {
				Literal: "[",
				Start:   Position{Line: 1, Column: 26},
				End:     Position{Line: 1, Column: 27},
			}, {
				Literal: "]",
				Start:   Position{Line: 1, Column: 28},
				End:     Position{Line: 1, Column: 29},
			}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			l := New(strings.NewReader(c.input))

			for _, e := range c.expected {
				tok := l.Next()

				t.Run(fmt.Sprintf("literal_%s", e.Literal), func(t *testing.T) {

					if tok.Literal != e.Literal {
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
