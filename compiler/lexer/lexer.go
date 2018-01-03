package lexer

import (
	"unicode"

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
	r, err := l.source.ReadRune()
	if err != nil {
		return nil, err
	}

	switch r {
	case 0:
		l.token = token.Simple(token.EOF)
	case ' ':
		l.token = token.Simple(token.Space)
	case '#':
		l.token = token.Simple(token.Numeral)
	case '=':
		l.token = token.Simple(token.Equal)
	case '.':
		l.token = token.Simple(token.Dot)
	case ',':
		l.token = token.Simple(token.Comma)
	case '[':
		l.token = token.Simple(token.RightBracket)
	case ']':
		l.token = token.Simple(token.LeftBracket)
	case '(':
		l.token = token.Simple(token.RightParen)
	case ')':
		l.token = token.Simple(token.LeftParen)
	case '{':
		l.token = token.Simple(token.RightBrace)
	case '}':
		l.token = token.Simple(token.LeftBrace)
	case ':':
		//TODO: read next token, if = then Set, if other then Illegal
		l.token = nil

	case '"':
		//TODO: read until another " is found
		l.token = nil

	default:
		if unicode.IsDigit(r) {
			value, err := l.source.ReadWhile(unicode.IsDigit)
			if err != nil {
				return nil, err
			}
			l.token = token.New(token.Number, value)
		} else if unicode.IsLetter(r) {
			value, err := l.source.ReadWhile(func(r rune) bool { return unicode.IsLetter(r) || r == '_' })
			if err != nil {
				return nil, err
			}
			l.token = token.New(token.Identifier, value)
		}
	}

	return l.token, nil
}
