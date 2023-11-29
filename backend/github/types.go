package github

type GithubRepository struct {
	Items []struct {
		Name  string `json:"name"`
		Owner struct {
			Login string `json:"login"`
		} `json:"owner"`
		DefaultBranch string `json:"default_branch"`
		Url           string `json:"html_url"`
	} `json:"items"`
}

type Project struct {
	Tree []Tree `json:"tree"`
}

type Tree struct {
	Path string `json:"path"`
}

type File struct {
	Name    string `json:"name"`
	Content string `json:"content"`
	Url     string `json:"html_url"`
}

type ParserFunctionCode interface {
	Parser(text []byte) ([]byte, error)
	FilterFile(trees []Tree) []string
	GetName() string
	GetExtension() string
}
