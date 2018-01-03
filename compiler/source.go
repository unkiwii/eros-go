package compiler

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Source struct {
	source      io.ReadCloser
	scanner     *bufio.Scanner
	runes       []rune
	currentRune rune
}

func (s *Source) Close() error {
	return s.source.Close()
}

func (s *Source) ReadRune() (r rune, err error) {
	if len(s.runes) == 0 {
		if ok := s.scanner.Scan(); !ok {
			if s.scanner.Err() != nil {
				return 0, fmt.Errorf("failed to scan next token: %s", s.scanner.Err())
			} else {
				return 0, nil
			}
		}
		s.runes = []rune(s.scanner.Text())
	}
	s.currentRune = s.runes[0]
	s.runes = s.runes[1:]
	return s.currentRune, nil
}

func (s *Source) ReadWhile(condition func(rune) bool) (value string, err error) {
	r := s.currentRune
	err = nil
	for condition(r) && err == nil {
		value += string(s.currentRune)
		r, err = s.ReadRune()
	}
	return value, err
}

func NewSourceFile(filename string) (*Source, error) {
	source, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return &Source{
		source:  source,
		scanner: newScanner(source),
	}, nil
}

func NewSource(str string) *Source {
	source := ioutil.NopCloser(strings.NewReader(str))
	return &Source{
		source:  source,
		scanner: newScanner(source),
	}
}

func newScanner(reader io.Reader) *bufio.Scanner {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanRunes)
	return scanner
}
