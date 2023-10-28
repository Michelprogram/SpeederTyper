package types

import (
	"golang.org/x/net/websocket"
)

type User struct {
	Ws       *websocket.Conn
	Username string `json:"username"`
}

func NewUser(ws *websocket.Conn, username string) *User {

	return &User{
		Ws:       ws,
		Username: username,
	}

}
