package send

import (
	"encoding/json"
	"github.com/michelprogram/speeder-typer/rooms"
	"github.com/michelprogram/speeder-typer/server"
	"github.com/michelprogram/speeder-typer/types"
	"golang.org/x/net/websocket"
)

var (
	ROOMS_INFO = "rooms-info"
	ROOM_INFO  = "room-info"
	USERNAME   = "username"
)

var _ server.EventSender = (*Sender)(nil)

type Sender struct{}

func NewSender() *Sender {
	return &Sender{}
}

func (s *Sender) RoomInfo(ws *websocket.Conn, game *rooms.Game) error {

	response := types.Response{
		Name: ROOM_INFO,
		Data: game,
	}

	res, err := json.Marshal(response)

	if err != nil {
		return err
	}

	for _, user := range game.Users {

		_, err = user.Ws.Write(res)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Sender) RoomsInfo(users map[*websocket.Conn]string, rooms map[string]*rooms.Game) error {

	stats := make([]types.RoomInfo, 0)

	var usernames []string

	for key, game := range rooms {
		stats = append(stats, types.RoomInfo{
			Name:    key,
			Players: game.GetUsers(),
		})
	}

	for _, username := range users {
		usernames = append(usernames, username)
	}

	response := types.Response{
		Name: ROOMS_INFO,
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
