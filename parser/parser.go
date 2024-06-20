package parser

import (
	"fmt"
	"strconv"

	"github.com/bousii/babbelas/ast"
	"github.com/bousii/babbelas/lexer"
)

type Parser struct {
	tokens []lexer.Token
}

func (p Parser) at() lexer.Token {
	return p.tokens[0]
}
func (p Parser) ProduceAST(sourceCode string) (ast.Program, error) {
	tokens, err := lexer.Tokenize(sourceCode)
	if err != nil {
		return ast.Program{}, err
	}
	program := ast.Program{}

	for i := 0; i < len(tokens); i++ {
		if tokens[i].Tokentype == lexer.EOF {
			break
		}
		stmt, err := p.parse_Stmt()
		if err != nil {
			return ast.Program{}, err
		}
		program.Body = append(program.Body, stmt)
	}
	return program, nil
}

func (p Parser) parse_Stmt() (ast.Stmt, error) {
	ret_Stmt, err := p.parse_Expr()
	if err != nil {
		return nil, err
	}
	return ret_Stmt, nil
}

func (p Parser) parse_Expr() (ast.Expr, error) {
	ret_Expr, err := p.parse_primary_Expr()
	if err != nil {
		return nil, err
	}
	return ret_Expr, nil
}

func (p Parser) parse_primary_Expr() (ast.Expr, error) {
	tokentype := p.at().Tokentype
	switch tokentype {
	case lexer.Identifier:
		return ast.Identifier{Symbol: p.at().Value}, nil
	case lexer.Number:
		num, _ := strconv.ParseFloat(p.at().Value, 64) // There should never be an error here, if there is, then the lexer is incorrect, I think...
		return ast.NumericLiteral{Value: num}, nil
	default:
		return nil, fmt.Errorf("Unidentified token identified in the parser %d", tokentype)
	}
}

// func (p Parser) parse_additive_Expr() (ast.Expr, error) {
// 	left := p.parse_primary_Expr()
// }
