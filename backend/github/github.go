package github

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/michelprogram/speeder-typer/server"
	"github.com/michelprogram/speeder-typer/types"
	"io"
	"math/rand"
	"net/http"
	"time"
)

var languages = []ParserFunctionCode{
	/*	Golang{},
	 */Typescript{},
	/*	Python{},
	 */ /*{"Jav", "js"},*/
}

var MAX_CHARACTER = 500

var PER_PAGE = 10

func pickAFile(files []string, owner, name string) (File, error) {
	var file File

	path := files[rand.Intn(len(files))]

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s", owner, name, path)
	err := getRequestGitHub(url, &file)
	if err != nil {
		return file, err
	}

	return file, nil

}

func getRequestGitHub(url string, data interface{}) error {
	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	//TODO : passer dans .env ou parametre de fct
	req.Header.Add("Authorization", "Bearer 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")

	response, err := client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	return nil
}

func FetchRandomCode(users []*types.User, sender server.EventSender) (codes string, err error) {

	var githubRepository GithubRepository
	var project Project
	var textInfo server.TextInfo

	rand.Seed(time.Now().UnixNano())

	language := languages[rand.Intn(len(languages))]

	textInfo.Language = language.GetName()

	go sender.CodeInfo(users, textInfo)

	url := fmt.Sprintf("https://api.github.com/search/repositories?q=language:%s&order=desc&per_page=%d", language.GetName(), PER_PAGE)
	err = getRequestGitHub(url, &githubRepository)

	if err != nil {
		return "", err
	}

	repo := githubRepository.Items[rand.Intn(PER_PAGE)]

	owner := repo.Owner.Login
	name := repo.Name
	branch := repo.DefaultBranch

	textInfo.RepoName = name
	textInfo.RepoUrl = repo.Url

	go sender.CodeInfo(users, textInfo)

	url = fmt.Sprintf("https://api.github.com/repos/%s/%s/git/trees/%s?recursive=1", owner, name, branch)
	err = getRequestGitHub(url, &project)
	if err != nil {
		return "", err
	}

	files := language.FilterFile(project.Tree)

	for {
		file, err := pickAFile(files, owner, name)
		if err != nil {
			return "", err
		}

		textInfo.FileName = file.Name
		textInfo.FullUrl = file.Url

		go sender.CodeInfo(users, textInfo)

		content, err := base64.StdEncoding.DecodeString(file.Content)
		if err != nil {
			return "", err
		}

		code, err := language.Parser(content)
		if err != nil {
			continue
		} else {
			return string(code), nil
		}
	}
}
