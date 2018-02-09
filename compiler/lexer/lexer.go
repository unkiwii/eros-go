package lexer

import (
	// "fmt"
	"io"
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
	r, _, err := l.source.ReadRune()
	if err != nil {
		if err == io.EOF {
			return token.Simple(token.EOF), nil
		}
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
		r, _, err = l.source.ReadRune()
		if err != nil {
			if err == io.EOF {
				return token.Simple(token.EOF), nil
			}
			return nil, err
		}
		if r == '=' {
			l.token = token.Simple(token.Set)
		}

	case '"':
		// l.source.Discard(1) // skip first '"'
		value, err := l.source.ReadWhile(isInsideString)
		if err != nil {
			if err == io.EOF {
				return token.Simple(token.EOF), nil
			}
			return nil, err
		}
		l.token = token.New(token.String, value)

	default:
		err = l.source.UnreadRune()
		if err != nil {
			return nil, err
		}
		if unicode.IsDigit(r) {
			l.token, err = l.readTokenWhile(isPartOfNumber, token.Number)
		} else if unicode.IsLetter(r) {
			l.token, err = l.readTokenWhile(isPartOfIdentifier, token.Identifier)
		}
	}

	return l.token, err
}

func isInsideString(r rune) bool {
	return r != '"'
}

func isPartOfNumber(r rune) bool {
	return unicode.IsDigit(r)
}

func isPartOfIdentifier(r rune) bool {
	return !unicode.IsSpace(r) && unicode.IsLetter(r)
}

func (l *Lexer) readTokenWhile(condition func(rune) bool, typ token.Type) (*token.Token, error) {
	value, err := l.source.ReadWhile(condition)
	if err != nil {
		if err == io.EOF {
			return token.Simple(token.EOF), nil
		}
		return nil, err
	}
	err = l.source.UnreadRune()
	if err != nil {
		return nil, err
	}
	return token.New(typ, value), nil
}
