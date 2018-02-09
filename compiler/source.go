package compiler

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Source struct {
	source io.ReadCloser
	reader *bufio.Reader
}

func (s *Source) Close() error {
	return s.source.Close()
}

func (s *Source) ReadRune() (r rune, size int, err error) {
	return s.reader.ReadRune()
}

func (s *Source) PeekRune() (r rune, size int, err error) {
	// TODO: implement
	return 0, 0, nil
}

func (s *Source) UnreadRune() error {
	return s.reader.UnreadRune()
}

func (s *Source) Discard(n int) (discarded int, err error) {
	return s.reader.Discard(n)
}

func (s *Source) ReadWhile(condition func(rune) bool) (str string, err error) {
	var buffer bytes.Buffer
	for r, _, err := s.ReadRune(); condition(r) && err == nil; r, _, err = s.ReadRune() {
		buffer.WriteRune(r)
	}
	return buffer.String(), err
}

func NewSourceFile(filename string) (*Source, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return &Source{
		source: file,
		reader: bufio.NewReader(file),
	}, nil
}

func NewSource(str string) *Source {
	r := ioutil.NopCloser(strings.NewReader(str))
	return &Source{
		source: r,
		reader: bufio.NewReader(r),
	}
}
