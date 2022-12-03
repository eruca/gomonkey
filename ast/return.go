package ast

import (
	"strings"

	"github.com/eruca/gomonkey/token"
)

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var builder strings.Builder
	builder.WriteString(rs.Token.Literal)
	builder.WriteByte(' ')
	builder.WriteString(rs.ReturnValue.String())

	return builder.String()
}
