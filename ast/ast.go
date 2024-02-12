package ast

import (
	"bytes"
	"monkey/token"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var writer bytes.Buffer
	for _, stmt := range p.Statements {
		writer.WriteString(stmt.String())
	}
	return writer.String()
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) String() string {
	var writer bytes.Buffer
	writer.WriteString(ls.TokenLiteral() + " ")
	writer.WriteString(ls.Name.String())
	writer.WriteString(" = ")

	if ls.Value != nil {
		writer.WriteString(ls.Value.String())
	}
	writer.WriteString(";")
	return writer.String()
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}

func (id *Identifier) expressionNode() {}

func (id *Identifier) TokenLiteral() string {
	return id.Token.Literal
}

func (id *Identifier) String() string {
	return id.Value
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) String() string {
	var writer bytes.Buffer
	writer.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		writer.WriteString(rs.ReturnValue.String())
	}
	writer.WriteString(";")
	return writer.String()
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) String() string       { return il.Token.Literal }
func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	var writer bytes.Buffer

	writer.WriteString("(")
	writer.WriteString(pe.Operator)
	writer.WriteString(pe.Right.String())
	writer.WriteString(")")

	return writer.String()
}

func (pe *PrefixExpression) expressionNode() {}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (pe *InfixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *InfixExpression) String() string {
	var writer bytes.Buffer

	writer.WriteString("(")
	writer.WriteString(pe.Left.String())
	writer.WriteString(" " + pe.Operator + " ")
	writer.WriteString(pe.Right.String())
	writer.WriteString(")")

	return writer.String()
}

func (pe *InfixExpression) expressionNode() {}

type Boolean struct {
	Token token.Token
	Value bool
}

func (pe *Boolean) TokenLiteral() string { return pe.Token.Literal }
func (pe *Boolean) String() string       { return pe.Token.Literal }
func (pe *Boolean) expressionNode()      {}

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var writer bytes.Buffer

	writer.WriteString("if")
	writer.WriteString(ie.Condition.String())
	writer.WriteString(" ")
	writer.WriteString(ie.Consequence.String())
	if ie.Alternative != nil {
		writer.WriteString("else ")
		writer.WriteString(ie.Alternative.String())
	}
	return writer.String()
}
func (ie *IfExpression) expressionNode() {}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var writer bytes.Buffer
	for _, s := range bs.Statements {
		writer.WriteString(s.String())
	}
	return writer.String()
}
func (bs *BlockStatement) expressionNode() {}

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var writer bytes.Buffer

	var params []string
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	writer.WriteString(fl.TokenLiteral())
	writer.WriteString("(")
	writer.WriteString(strings.Join(params, ", "))
	writer.WriteString(") ")
	writer.WriteString(fl.Body.String())

	return writer.String()
}
func (fl *FunctionLiteral) expressionNode() {}

type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var writer bytes.Buffer

	var args []string
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	writer.WriteString(ce.Function.String())
	writer.WriteString("(")
	writer.WriteString(strings.Join(args, ", "))
	writer.WriteString(")")

	return writer.String()
}
func (ce *CallExpression) expressionNode() {}

type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) String() string       { return sl.Token.Literal }
func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
