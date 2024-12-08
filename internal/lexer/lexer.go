package lexer

import (
	"fmt"
	"io"
	"strings"
	"text/scanner"
	"unicode"

	"github.com/SimonRichardson/runscript/internal/token"
)

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

func (l *Lexer) Next() token.Token {
	start := l.scanner.Pos()
	for tok := l.scanner.Scan(); tok != scanner.EOF; tok = l.scanner.Scan() {
		switch tok {
		case '\n', '\t', ' ':
			start = l.scanner.Pos()
			continue

		default:
			return l.makeToken(start)
		}
	}

	end := l.scanner.Pos()
	return token.Token{
		Type: token.EOF,
		Start: token.Position{
			Line:   start.Line,
			Column: start.Column - 1,
		},
		End: token.Position{
			Line:   end.Line,
			Column: end.Column - 1,
		},
	}
}

func (l *Lexer) makeToken(start scanner.Position) token.Token {
	literal := l.scanner.TokenText()
	end := l.scanner.Pos()
	return token.Token{
		Type: l.tokenTypeFromString(literal),
		Start: token.Position{
			Line:   start.Line,
			Column: start.Column - 1,
		},
		End: token.Position{
			Line:   end.Line,
			Column: end.Column - 1,
		},
		Literal: literal,
	}
}

func (l *Lexer) tokenTypeFromString(literal string) token.TokenType {
	switch literal {
	case "=":
		return token.ASSIGN
	case ";":
		return token.SEMICOLON
	case "(":
		return token.LPAREN
	case ")":
		return token.RPAREN
	case ",":
		return token.COMMA
	case "+":
		return token.PLUS
	case "-":
		return token.MINUS
	case "{":
		return token.LBRACE
	case "}":
		return token.RBRACE
	case "[":
		return token.LBRACKET
	case "]":
		return token.RBRACKET
	case ":":
		return token.COLON
	case ".":
		return token.PERIOD
	case "==":
		return token.EQ
	case "!=":
		return token.NEQ
	case "<":
		return token.LT
	case ">":
		return token.GT
	case "!":
		return token.BANG
	case "*":
		return token.ASTERISK
	case "/":
		return token.SLASH
	case "%":
		return token.MOD
	case "&":
		return token.BITAND
	case "|":
		return token.BITOR
	case "&&":
		return token.CONDAND
	case "||":
		return token.CONDOR
	case "<=":
		return token.LE
	case ">=":
		return token.GE
	case "=>":
		return token.LAMBDA
	case "_":
		return token.UNDERSCORE
	default:
		fmt.Println(literal)
		if unicode.IsDigit(rune(literal[0])) {
			if strings.Contains(literal, ".") {
				return token.FLOAT
			}
			return token.INT
		} else if len(literal) > 1 && rune(literal[0]) == '.' && unicode.IsDigit(rune(literal[1])) {
			return token.FLOAT
		}
		return token.IDENT
	}
}
