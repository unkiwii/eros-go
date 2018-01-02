package lexer

import (
	"fmt"

	"github.com/unkiwii/eros-go/compiler"
	"github.com/unkiwii/eros-go/compiler/token"
)

type Lexer struct {
	source *compiler.Source
	token  *token.Token
}

func NewLexer(source *compiler.Source) *Lexer {
	return &Lexer{source, nil}
}

func (l *Lexer) Close() error {
	return l.source.Close()
}

func (l *Lexer) NextToken() (*token.Token, error) {
	txt, size, err := l.source.Read()
	if err != nil {
		return nil, err
	}

	fmt.Printf("read: '%s' (%d)\n", txt, size)

	switch txt {
	case " ":
		l.token = token.Simple(token.Space)
	case "#":
		l.token = token.Simple(token.Numeral)
	case "=":
		l.token = token.Simple(token.Equal)
	case ".":
		l.token = token.Simple(token.Dot)
	case ",":
		l.token = token.Simple(token.Comma)
	case "[":
		l.token = token.Simple(token.RightBracket)
	case "]":
		l.token = token.Simple(token.LeftBracket)
	case "(":
		l.token = token.Simple(token.RightParen)
	case ")":
		l.token = token.Simple(token.LeftParen)
	case "{":
		l.token = token.Simple(token.RightBrace)
	case "}":
		l.token = token.Simple(token.LeftBrace)
	case ":":
		l.token = token.Simple(token.Comma)

	default:
		// if is number
		//   read number
		// else if is alpha
		//   read identifier
	}

	return l.token, nil
}
