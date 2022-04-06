package ast

import "github.com/SerhiiCho/parser/token"

type LetStatement struct {
	Token token.Token // the token.Let token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
