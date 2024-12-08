package lexer

import (
	"io"
	"text/scanner"
)

type Position struct {
	Line   int
	Column int
}

type Token struct {
	Start      Position
	End        Position
	Literal    string
	Terminated bool
}

type Lexer struct {
	scanner *scanner.Scanner
}

func New(reader io.Reader) *Lexer {
	scanner := new(scanner.Scanner)

	scanner.Init(reader)
	scanner.Whitespace ^= 1<<'\n' | 1<<'\t' | 1<<' '

	return &Lexer{
		scanner: scanner,
	}
}

func (l *Lexer) Next() Token {
	start := l.scanner.Pos()
	for tok := l.scanner.Scan(); tok != scanner.EOF; tok = l.scanner.Scan() {
		literal := l.scanner.TokenText()
		switch tok {
		case '\n', '\t', ' ':
			start = l.scanner.Pos()
			continue

		default:
			end := l.scanner.Pos()
			return Token{
				Start: Position{
					Line:   start.Line,
					Column: start.Column - 1,
				},
				End: Position{
					Line:   end.Line,
					Column: end.Column - 1,
				},
				Literal: literal,
			}
		}
	}

	end := l.scanner.Pos()
	return Token{
		Start: Position{
			Line:   start.Line,
			Column: start.Column - 1,
		},
		End: Position{
			Line:   end.Line,
			Column: end.Column - 1,
		},
		Terminated: true,
	}
}
