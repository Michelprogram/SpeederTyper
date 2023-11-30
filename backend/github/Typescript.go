package github

import (
	"bytes"
	"errors"
	"math/rand"
	"regexp"
	"strings"
)

type Typescript struct {
}

var _ ParserFunctionCode = (*Typescript)(nil)

func (g Typescript) FilterFile(trees []Tree) []string {
	var res []string
	extension := ".ts"
	cypress := ".cy.ts"
	test := ".test.ts"

	for _, tree := range trees {

		if strings.HasSuffix(tree.Path, extension) && !strings.HasSuffix(tree.Path, cypress) && !strings.HasSuffix(tree.Path, test) {
			res = append(res, tree.Path)
		}
	}

	return res
}

func (g Typescript) GetName() string {
	return "Typescript"
}

func (g Typescript) GetExtension() string {
	return ".ts"
}

func (g Typescript) Parser(text []byte) ([]byte, error) {

	regex := regexp.MustCompile("function ")

	matches := regex.FindAllIndex(text, -1)

	if len(matches) == 0 {
		return nil, errors.New("No func found in file")
	}

	startIndex := matches[rand.Intn(len(matches))][0]

	text = text[startIndex:]

	bracketCounter := 1
	i := bytes.IndexByte(text, '{') + 1

	for bracketCounter > 0 && i < len(text) {

		letter := text[i]

		if letter == '{' {
			bracketCounter++
		} else if letter == '}' {
			bracketCounter--
		}

		i++
	}

	if len(text[:1]) > MAX_CHARACTER {
		return nil, errors.New("Text too long for a game.")
	}

	return []byte("function"), nil

	return text[:i], nil
}
