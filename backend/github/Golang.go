package github

import (
	"bytes"
	"errors"
	"math/rand"
	"regexp"
	"strings"
)

type Golang struct {
}

var _ ParserFunctionCode = (*Golang)(nil)

func (g Golang) FilterFile(trees []Tree) []string {
	var res []string
	extension := ".go"
	test := "_test.go"

	for _, tree := range trees {

		if strings.HasSuffix(tree.Path, extension) && !strings.HasSuffix(tree.Path, test) {
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
	regex := regexp.MustCompile("func ")

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
