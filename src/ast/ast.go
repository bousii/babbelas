package ast

type NodeType string

const (
	program        NodeType = "Program"
	numericLiteral NodeType = "NumericLiteral"
	identifier     NodeType = "Identifier"
	binaryExpr     NodeType = "BinaryExpr"
)

type Stmt interface {
	GetKind() NodeType
}

type Expr interface {
}

type Program struct {
	Body []Stmt
}

func (p Program) GetKind() NodeType {
	return program
}

type NumericLiteral struct {
	Value int
}

func (n NumericLiteral) GetKind() NodeType {
	return numericLiteral
}

type Identifier struct {
	symbol string
}

func (i Identifier) GetKind() NodeType {
	return identifier
}

type BinaryExpr struct {
	left     Expr
	right    Expr
	operator string
}

func (b BinaryExpr) GetKind() NodeType {
	return binaryExpr
}
