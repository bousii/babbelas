package ast

type NodeType string

const (
	program        NodeType = "Program"
	numericLiteral NodeType = "NumericLiteral"
	identifier     NodeType = "Identifier"
	binaryExpr     NodeType = "BinaryExpr"
	expr           NodeType = "Expression"
)

type Stmt interface {
	GetKind() NodeType
}

type Expr interface {
	GetKind() NodeType
}

type Program struct {
	Body []Stmt
}

func (p Program) GetKind() NodeType {
	return program
}

type NumericLiteral struct {
	Value float64
}

func (n NumericLiteral) GetKind() NodeType {
	return numericLiteral
}

type Identifier struct {
	Symbol string
}

func (i Identifier) GetKind() NodeType {
	return identifier
}

type BinaryExpr struct {
	Left     Expr
	Right    Expr
	Operator string
}

func (b BinaryExpr) GetKind() NodeType {
	return binaryExpr
}
