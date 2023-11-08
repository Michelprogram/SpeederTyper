package types

import (
	"golang.org/x/net/websocket"
)

type User struct {
	Ws       *websocket.Conn `json:"-"`
	Username string          `json:"username"`
	IsReady  bool            `json:"isReady"`
	Position int             `json:"position"`
}

func NewUser(ws *websocket.Conn, username string) *User {

	return &User{
		Ws:       ws,
		Username: username,
		IsReady:  false,
		Position: 0,
	}

}

func (u *User) SetIsReady(isready bool) {
	u.IsReady = isready
}

type ByPosition []*User

func (a ByPosition) Len() int           { return len(a) }
func (a ByPosition) Less(i, j int) bool { return a[i].Position > a[j].Position }
func (a ByPosition) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
