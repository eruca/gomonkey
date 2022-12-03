package ast

import (
	"strings"

	"github.com/eruca/gomonkey/token"
)

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out strings.Builder

	out.WriteByte('(')
	out.WriteString(ie.Left.String())
	out.WriteByte(' ')
	out.WriteString(ie.Operator)
	out.WriteByte(' ')
	out.WriteString(ie.Right.String())
	out.WriteByte(')')

	return out.String()
}
