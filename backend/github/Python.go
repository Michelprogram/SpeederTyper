package github

import (
	"bytes"
	"errors"
	"log"
	"math/rand"
	"regexp"
	"strings"
)

type Python struct {
}

var _ ParserFunctionCode = (*Python)(nil)

func (p Python) Parser(text []byte) ([]byte, error) {
	regex := regexp.MustCompile("def\\s+([a-zA-Z_][a-zA-Z0-9_]*)\\s*\\([^)]*\\)\\s*:\n")

	matches := regex.FindAllIndex(text, -1)

	if len(matches) == 0 {
		return nil, errors.New("No def found in file")
	}

	startIndex := matches[rand.Intn(len(matches))][0]

	text = text[startIndex:]

	i := bytes.IndexByte(text, ':')

	for i < len(text) {

		end := bytes.IndexByte(text[i:], '\n')

		line := text[i : i+end]

		if !bytes.HasSuffix(line, []byte("\t")) {
			break
		}

		i += end + 1

	}

	log.Println(string(text[:i]))

	panic("oui")

	if len(text[:i]) > MAX_CHARACTER {
		return nil, errors.New("Text too long for a game.")
	}

	return text[:i], nil
}

func (p Python) FilterFile(trees []Tree) []string {
	var res []string
	init := "__init"
	test := "_test.py"

	for _, tree := range trees {
		if strings.HasSuffix(tree.Path, p.GetExtension()) &&
			!strings.HasPrefix(tree.Path, init) &&
			!strings.HasSuffix(tree.Path, test) {
			res = append(res, tree.Path)
		}
	}

	return res
}

func (p Python) GetName() string {
	return "Python"
}

func (p Python) GetExtension() string {
	return ".py"
}
