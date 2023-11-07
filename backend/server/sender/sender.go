package sender

import (
	"encoding/json"
	"github.com/michelprogram/speeder-typer/rooms"
	"github.com/michelprogram/speeder-typer/server"
	"github.com/michelprogram/speeder-typer/types"
	"golang.org/x/net/websocket"
)

var (
	GAME_CANT_START = "game-cant-start"
	GAME_END        = "game-end"

	ROOMS_INFO            = "rooms-info"
	ROOM_INFO             = "room-info"
	ROOM_CREATED          = "room-created"
	ROOM_JOIN_BY_USERNAME = "room-join-by-username"

	USERNAME = "username"
	PONG     = "pong"
)

var _ server.EventSender = (*Sender)(nil)

type Sender struct{}

func NewSender() *Sender {
	return &Sender{}
}

func (s *Sender) JoinedRoomByUsername(conn *websocket.Conn, id string) error {

	response := types.Response{
		Name: ROOM_JOIN_BY_USERNAME,
	}

	if id == "" {
		response.Data = "-1"
	} else {
		response.Data = id
	}

	return s.JsonResponse(conn, response)

}

func (s *Sender) RoomCreated(conn *websocket.Conn, id string) error {
	response := types.Response{
		Name: ROOM_CREATED,
		Data: id,
	}

	return s.JsonResponse(conn, response)
}

func (s *Sender) Pong(conn *websocket.Conn) error {
	response := types.Response{
		Name: PONG,
		Data: "Pong",
	}

	return s.JsonResponse(conn, response)
}

func (s *Sender) GameCantStart(conn *websocket.Conn) error {
	response := types.Response{
		Name: GAME_CANT_START,
	}

	return s.JsonResponse(conn, response)

}

func (s *Sender) GameEnd(game *rooms.Game) error {

	//TODO : faire le classement

	response := types.Response{
		Name: GAME_END,
		Data: game.Users,
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

func (s *Sender) RoomInfo(game *rooms.Game) error {

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

	stats := make([]*types.RoomInfo, 0)

	var usernames []string

	for key, game := range rooms {
		stats = append(stats, types.NewRoomInfo(key, game.Status, game.GetUsers()))
	}

	for _, username := range users {
		usernames = append(usernames, username)
	}

	response := types.Response{
		Name: ROOMS_INFO,
		Data: NewStats(stats, usernames),
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

	return err
}

func (s *Sender) UserLogged(ws *websocket.Conn, username string) error {

	response := types.Response{
		Name: USERNAME,
		Data: username,
	}

	return s.JsonResponse(ws, response)
}
