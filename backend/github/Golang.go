package github

import (
	"bytes"
	"errors"
	"math/rand"
	"regexp"
)

type Golang struct {
}

var _ ParserFunctionCode = (*Golang)(nil)

func (g Golang) FilterFile(trees []Tree) []string {
	var res []string
	extension := ".go"
	extensionTestFile := "_test.go"

	sizeExtension := len(extension)
	sizeTestFile := len(extensionTestFile)

	for _, tree := range trees {

		if tree.Path[len(tree.Path)-sizeExtension:] == extension && tree.Path[len(tree.Path)-sizeTestFile:] != extensionTestFile {
			res = append(res, tree.Path)
		}
	}

	return res
}

func (g Golang) GetName() string {
	return "Go"
}

func (g Golang) GetExtension() string {
	return ".go"
}

func (g Golang) Parser(text []byte) ([]byte, error) {
	regex := regexp.MustCompile("func")

	matches := regex.FindAllIndex(text, -1)

	if len(matches) == 0 {
		return nil, errors.New("No func found in file")
	}

	startIndex := matches[rand.Intn(len(matches))][0]

	text = text[startIndex:]

	bracketCounter := 1
	i := bytes.IndexByte(text, '{') + 1

	for bracketCounter > 0 {

		letter := text[i]

		if letter == '{' {
			bracketCounter++
		} else if letter == '}' {
			bracketCounter--
		}

		i++
	}

	return text[:i], nil
}
