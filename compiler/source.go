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
	source  io.ReadCloser
	scanner *bufio.Scanner
}

func (s *Source) Close() error {
	return s.source.Close()
}

func (s *Source) Read() (txt string, size int, err error) {
	if ok := s.scanner.Scan(); !ok {
		return "", 0, fmt.Errorf("failed to scan next token: %s", s.scanner.Err())
	}
	//TODO: check this
	txt = s.scanner.Text()
	return txt, len(txt), nil
}

func NewSourceFile(filename string) (*Source, error) {
	source, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return &Source{source, newScanner(source)}, nil
}

func NewSource(str string) *Source {
	source := ioutil.NopCloser(strings.NewReader(str))
	return &Source{source, newScanner(source)}
}

func newScanner(reader io.Reader) *bufio.Scanner {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanRunes)
	return scanner
}
