package rooms

import (
	"github.com/michelprogram/speeder-typer/types"
	"log"
	"time"
)

const (
	Waiting int = 0
	Gaming  int = 1
	Finish  int = 3
)

type Game struct {
	Users     []*types.User `json:"users"`
	CreatedBy string        `json:"created_by"`
	Text      string        `json:"text"`
	Id        string        `json:"id"`
	Status    int           `json:"status"`
	StartAt   time.Time     `json:"-"`
	EndAt     time.Time     `json:"-"`
}

func NewGame(id string, username string) *Game {

	return &Game{
		Users:     make([]*types.User, 0),
		CreatedBy: username,
		Text:      "",
		Id:        id,
		Status:    Waiting,
	}
}

func (g *Game) UserJoin(user *types.User) {
	user.SetIsReady(false)
	user.SetPosition(0)
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

	log.Println(len(g.Users))

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

	return count == len(g.Users)-1
}

func (g *Game) SetStatus(status int) {
	g.Status = status
}

func (g *Game) SetUserPosition(username string) {

	for _, user := range g.Users {
		if user.Username == username {
			user.Position++
		}
	}

}

func (g Game) FindUserByUsername(username string) *types.User {
	for _, user := range g.Users {
		if user.Username == username {
			return user
		}
	}

	return nil
}

func (g *Game) SetStartAt(time time.Time) {
	g.StartAt = time
}

func (g *Game) SetEndAt(time time.Time) {
	g.EndAt = time
}

func (g *Game) SetText(text string) {
	g.Text = text
}
