package test

import (
	"github.com/android-test-runner/atr/files"
	"strings"
)

type Parser interface {
	ParseFromFile(path string) ([]Test, error)
	Parse(testsUnparsed []string) []Test
}

type parserImpl struct {
	files files.Files
}

func NewParser() Parser {
	return parserImpl{
		files: files.New(),
	}
}

func (parser parserImpl) ParseFromFile(path string) ([]Test, error) {
	lines, err := parser.files.ReadLines(path)
	if err != nil {
		return nil, err
	}
	return parser.Parse(lines), nil
}

func (parserImpl) Parse(testsUnparsed []string) []Test {
	var tests []Test
	for _, testUnparsed := range testsUnparsed {
		tests = append(tests, parseTest(testUnparsed))
	}

	return tests
}

func parseTest(testUnparsed string) Test {
	tokens := strings.Split(testUnparsed, "#")
	return Test{
		Class:  tokens[0],
		Method: tokens[1],
	}
}
