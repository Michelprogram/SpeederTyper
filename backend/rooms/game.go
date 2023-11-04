package rooms

import (
	"github.com/michelprogram/speeder-typer/types"
	"github.com/michelprogram/speeder-typer/utils"
)

type Game struct {
	Users     []*types.User `json:"users"`
	CreatedBy string        `json:"created_by"`
	Text      string        `json:"text"`
	Id        string        `json:"id"`
	Status    bool          `json:"status"`
}

func NewGame(id string, username string) (*Game, error) {

	text, err := utils.FetchRandomText()

	if err != nil {
		return nil, err
	}

	return &Game{
		Users:     make([]*types.User, 0),
		CreatedBy: username,
		Text:      text,
		Id:        id,
		Status:    false,
	}, nil
}

func (g *Game) UserJoin(user *types.User) {
	g.Users = append(g.Users, user)
}

func (g *Game) UserLeft(username string) int {
	for index, user := range g.Users {
		if user.Username == username {
			g.Users = append(g.Users[:index], g.Users[index+1:]...)
			break
		}
	}
	return len(g.Users)
}

func (g *Game) Usernames() (usernames []string) {

	for _, user := range g.Users {
		usernames = append(usernames, user.Username)
	}

	return usernames
}

func (g Game) GetUsers() []*types.User {
	return g.Users
}

func (g *Game) IsUserInGame(username string) *types.User {
	for _, user := range g.Users {
		if user.Username == username {
			return user
		}
	}

	return nil
}

func (g Game) SetUserReady(username string, ready bool) {
	for _, user := range g.Users {
		if user.Username == username {
			user.SetIsReady(ready)
		}
	}

}

func (g Game) IsEveryoneReady() bool {

	count := 0

	for _, user := range g.Users {
		if user.IsReady {
			count++
		}
	}

	//Createur of user can't be ready
	return count == len(g.Users)-1
}
