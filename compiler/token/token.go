package token

import (
	"fmt"
)

type Type uint8

const (
	EOF Type = iota
	Illegal

	Space
	Numeral
	Equal
	Set
	Dot
	Comma
	LeftBracket
	RightBracket
	LeftParen
	RightParen
	LeftBrace
	RightBrace
	Number
	Identifier
	String
)

type Token struct {
	typ Type
	val string
}

var simples = map[Type]*Token{
	Space:        &Token{Space, " "},
	Numeral:      &Token{Numeral, "#"},
	Equal:        &Token{Equal, "="},
	Set:          &Token{Set, ":="},
	Dot:          &Token{Dot, "."},
	Comma:        &Token{Comma, ","},
	LeftBracket:  &Token{LeftBracket, "["},
	RightBracket: &Token{RightBracket, "]"},
	LeftParen:    &Token{LeftParen, "("},
	RightParen:   &Token{RightParen, ")"},
	LeftBrace:    &Token{LeftBrace, "{"},
	RightBrace:   &Token{RightBrace, "}"},
}

func (t Type) String() string {
	switch t {
	case Illegal:
		return "<illegal>"
	case Number:
		return "<number>"
	case Identifier:
		return "<id>"
	case String:
		return "<string>"
	default:
		return fmt.Sprintf("<%s>", simples[t].val)
	}
}

func (t *Token) Type() Type {
	if t == nil {
		return Illegal
	}
	return t.typ
}

func (t *Token) Value() string {
	if t == nil {
		return "<illegal>"
	}
	return t.val
}

func Simple(t Type) *Token {
	return simples[t]
}

func New(t Type, v string) *Token {
	return &Token{t, v}
}
