package ast

import (
	"monkey/token"
	"testing"
)

func TestProgram_String(t *testing.T) {
	p := Program{
		[]Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}
	if p.String() != "let myVar = anotherVar" {
		t.Errorf("p.String() wrong. got=%q", p.String())
	}
}
