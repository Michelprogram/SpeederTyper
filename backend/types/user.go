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
