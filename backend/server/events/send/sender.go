package send

import (
	"encoding/json"
	"github.com/michelprogram/speeder-typer/server"
	"github.com/michelprogram/speeder-typer/types"
	"golang.org/x/net/websocket"
)

var (
	ROOM_INFO = "room-info"
	USERNAME  = "username"
)

var _ server.EventSender = (*Sender)(nil)

type Sender struct{}

func NewSender() *Sender {
	return &Sender{}
}

func (s *Sender) RoomsInfo(users map[*websocket.Conn]string, rooms map[string][]*types.User) error {

	stats := make([]types.RoomInfo, 0)

	usernames := make([]string, len(users))

	for key, value := range rooms {
		stats = append(stats, types.RoomInfo{
			Name:    key,
			Players: value,
		})
	}

	for _, username := range users {
		usernames = append(usernames, username)
	}

	response := types.Response{
		Name: ROOM_INFO,
		Data: struct {
			Stats []types.RoomInfo `json:"stats"`
			Users []string         `json:"users"`
		}{stats, usernames},
	}

	res, err := json.Marshal(response)

	if err != nil {
		return err
	}

	for user := range users {
		user.Write(res)
	}

	return nil
}

func (s *Sender) JsonResponse(ws *websocket.Conn, response types.Response) error {

	res, err := json.Marshal(response)

	if err != nil {
		return err
	}

	_, err = ws.Write(res)

	if err != nil {
		return err
	}

	return nil
}

func (s *Sender) UserLogged(ws *websocket.Conn, username string) error {

	response := types.Response{
		Name: USERNAME,
		Data: username,
	}

	return s.JsonResponse(ws, response)
}
