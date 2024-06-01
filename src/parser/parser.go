package parser

import (
	"github.com/bousii/babbelas/src/ast"
	"github.com/bousii/babbelas/src/lexer"
)

type Parser struct {
	tokens []lexer.Token
}

func (p Parser) produceAST(sourceCode string) (ast.Program, error) {
	tokens, err := lexer.Tokenize(sourceCode)
	if err != nil {
		return ast.Program{}, err
	}
	program := ast.Program{}
	return program, nil
}
