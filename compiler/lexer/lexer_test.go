package lexer

import (
	"testing"

	"github.com/unkiwii/eros-go/compiler"
	"github.com/unkiwii/eros-go/compiler/token"
)

func TestEmpty(t *testing.T) {
	testLexerCase(t, "",
		token.Simple(token.EOF))
}

func TestNumber(t *testing.T) {
	testLexerCase(t, "1234567890",
		token.New(token.Number, "1234567890"),
		token.Simple(token.EOF))

	//TODO: add number tests (decimals, complex, fractionals)
}

func TestIdentifier(t *testing.T) {
	testLexerCase(t, "qwertyuiopasdfghjklzxcvbnm",
		token.New(token.Identifier, "qwertyuiopasdfghjklzxcvbnm"),
		token.Simple(token.EOF))

	testLexerCase(t, "QWERTYUIOPASDFGHJKLZXCVBNM",
		token.New(token.Identifier, "QWERTYUIOPASDFGHJKLZXCVBNM"),
		token.Simple(token.EOF))

	//TODO: add identifier tests (decimals, complex, fractionals)
}

func TestString(t *testing.T) {
	testLexerCase(t, "\"string\"",
		token.New(token.String, "string"),
		token.Simple(token.EOF))

	testLexerCase(t, "\"\"",
		token.New(token.String, ""),
		token.Simple(token.EOF))

	//TODO: add string tests (quotes inside quotes, simple, double)
}

func TestEqualSetCombination(t *testing.T) {
	testLexerCase(t, "==:=:==",
		token.Simple(token.Equal),
		token.Simple(token.Equal),
		token.Simple(token.Set),
		token.Simple(token.Set),
		token.Simple(token.Equal),
		token.Simple(token.EOF))
}

func TestAssignments(t *testing.T) {
	testLexerCase(t, "Number n := 123.",
		token.New(token.Identifier, "Number"),
		token.Simple(token.Space),
		token.New(token.Identifier, "n"),
		token.Simple(token.Space),
		token.Simple(token.Set),
		token.Simple(token.Space),
		token.New(token.Number, "123"),
		token.Simple(token.Dot),
		token.Simple(token.EOF))

	testLexerCase(t, "n := 456.",
		token.New(token.Identifier, "n"),
		token.Simple(token.Space),
		token.Simple(token.Set),
		token.Simple(token.Space),
		token.New(token.String, "456"),
		token.Simple(token.Dot),
		token.Simple(token.EOF))

	testLexerCase(t, "String s := \"hello\".",
		token.New(token.Identifier, "String"),
		token.Simple(token.Space),
		token.New(token.Identifier, "s"),
		token.Simple(token.Space),
		token.Simple(token.Set),
		token.Simple(token.Space),
		token.New(token.String, "hello"),
		token.Simple(token.Dot),
		token.Simple(token.EOF))

	testLexerCase(t, "s := \"hello\".",
		token.New(token.Identifier, "s"),
		token.Simple(token.Space),
		token.Simple(token.Set),
		token.Simple(token.Space),
		token.New(token.String, "hello"),
		token.Simple(token.Dot),
		token.Simple(token.EOF))
}

func testLexerCase(t *testing.T, input string, tokens ...*token.Token) {
	lexer := NewLexer(compiler.NewSource(input))
	var got *token.Token
	var err error
	for _, expected := range tokens {
		if got, err = lexer.NextToken(); err != nil {
			t.Error("error with lexer.NextToken()", err)
		}
		if expected.Type() != got.Type() || expected.Value() != got.Value() {
			t.Errorf("testing: '%s', expected '%s' %s, but got '%s' %s", input, expected.Value(), expected.Type(), got.Value(), got.Type())
		}
	}
}
