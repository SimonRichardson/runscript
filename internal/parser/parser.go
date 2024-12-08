package parser

import (
	"github.com/SimonRichardson/runscript/internal/ast"
	"github.com/SimonRichardson/runscript/internal/token"
)

const (
	LOWEST = iota
	PCONDOR
	PCONDAND
	EQUALS
	LESSGREATER
	PPRODUCT
	CALL
	INDEX
)

var precedence = map[token.TokenType]int{
	token.CONDOR:   PCONDOR,
	token.CONDAND:  PCONDAND,
	token.EQ:       EQUALS,
	token.NEQ:      EQUALS,
	token.LPAREN:   CALL,
	token.LAMBDA:   CALL,
	token.LT:       LESSGREATER,
	token.LE:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.GE:       LESSGREATER,
	token.LBRACKET: INDEX,
	token.PERIOD:   INDEX,
}

type Lexer interface {
	Next() token.Token
}

type PrefixFunc func() (ast.Expression, error)
type InfixFunc func(ast.Expression) (ast.Expression, error)

type Parser struct {
	lexer Lexer

	prefix map[token.TokenType]PrefixFunc
	infix  map[token.TokenType]InfixFunc

	prevToken    token.Token
	currentToken token.Token
	peekToken    token.Token
}

func New(lexer Lexer) *Parser {
	p := &Parser{
		lexer: lexer,
	}

	p.prefix = map[token.TokenType]PrefixFunc{
		token.IDENT: p.parseIdentifier,
	}
	p.infix = map[token.TokenType]InfixFunc{}

	return p
}

func (p *Parser) Parse() (ast.QueryExpression, error) {
	var exp ast.QueryExpression
	for p.currentToken.Type != token.EOF {
		expr, err := p.parseExpressionStatement()
		if err != nil {
			return exp, err
		}
		exp.Expressions = append(exp.Expressions, expr)
		p.nextToken()
	}
	return exp, nil
}

func (p *Parser) nextToken() {
	p.prevToken = p.currentToken
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.Next()
}

func (p *Parser) parseExpressionStatement() (ast.Expression, error) {
	stmt := ast.ExpressionStatement{
		Token: p.currentToken,
	}
	expr, err := p.parseExpression(LOWEST)
	if err != nil {
		return nil, err
	}
	stmt.Expression = expr

	if p.isPeekToken(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt, nil
}

func (p *Parser) parseIdentifier() (ast.Expression, error) {
	return ast.Identifier{
		Token: p.currentToken,
	}, nil
}

func (p *Parser) parseExpression(precedence int) (ast.Expression, error) {
	prefix := p.prefix[p.currentToken.Type]
	if prefix == nil {
		if p.currentToken.Type != token.EOF {
			return nil, SyntaxError{
				TokenType: p.currentToken.Type,
				Positions: []token.Position{
					p.currentToken.Start,
					p.currentToken.End,
				},
			}
		}
		return nil, nil
	}
	leftExp, err := prefix()
	if err != nil {
		return nil, err
	}
	// Run the infix function until the next token has
	// a higher precedence.
	for !p.isPeekToken(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infix[p.peekToken.Type]
		if infix == nil {
			return leftExp, nil
		}
		p.nextToken()
		if leftExp, err = infix(leftExp); err != nil {
			return nil, err
		}
	}

	return leftExp, nil
}

func (p *Parser) isPeekToken(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedence[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}
