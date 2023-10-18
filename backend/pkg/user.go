package pkg

import "golang.org/x/net/websocket"

type User struct {
	Ws       *websocket.Conn
	Username string `json:"username"`
}

func NewUser(ws *websocket.Conn, username string) *User {

	response := Response{
		Name: USERNAME,
		Data: username,
	}

	SendJsonResponse(ws, response)

	return &User{
		Ws:       ws,
		Username: username,
	}

}
