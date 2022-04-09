package ast

import "github.com/SerhiiCho/parser/token"

type LetStatement struct {
	Token token.Token // {Type:"LET", Literal:"let"}
	Name  *Identifier // {Token:token.Token{Type:"IDENT", Literal:"x"}, Value:"x"}
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
