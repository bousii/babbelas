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
		stmt, err := parse_Stmt(tokens[i])
		if err != nil {
			return ast.Program{}, err
		}
		program.Body = append(program.Body, stmt)
	}
	return program, nil
}

func parse_Stmt(token lexer.Token) (ast.Stmt, error) {
	ret_Stmt, err := parse_Expr(token)
	if err != nil {
		return nil, err
	}
	return ret_Stmt, nil
}

func parse_Expr(token lexer.Token) (ast.Expr, error) {
	ret_Expr, err := parse_first_Expr(token)
	if err != nil {
		return nil, err
	}
	return ret_Expr, nil
}

func parse_first_Expr(token lexer.Token) (ast.Expr, error) {
	tokentype := token.Tokentype
	switch tokentype {
	case lexer.Identifier:
		return ast.Identifier{Symbol: token.Value}, nil
	case lexer.Number:
		num, _ := strconv.ParseFloat(token.Value, 64) // There should never be an error here, if there is, then the lexer is incorrect, I think...
		return ast.NumericLiteral{Value: num}, nil
	default:
		return nil, fmt.Errorf("Unidentified token identified in the parser %d", tokentype)
	}
}
