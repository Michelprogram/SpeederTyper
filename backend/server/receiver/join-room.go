package receiver

import (
	"encoding/json"
	"github.com/michelprogram/speeder-typer/server"
	"github.com/michelprogram/speeder-typer/types"
	"golang.org/x/net/websocket"
	"log"
)

type JoinRoom struct{}

func (j JoinRoom) HandleData(s *server.Server, ws *websocket.Conn, data json.RawMessage) error {

	var idRoom JSONIdRoom

	err := json.Unmarshal(data, &idRoom)

	if err != nil {
		return err
	}

	defer s.Sender.RoomsInfo(s.Users, s.Rooms.Games)

	username := s.Users[ws]

	k, v := s.Rooms.Games[idRoom.Id]

	log.Println(k, v)

	if s.Rooms.IsUserInRoom(username, idRoom.Id) != nil {
		return nil
	}

	user := types.NewUser(ws, username)

	s.Rooms.JoinRoom(user, idRoom.Id)

	game := s.Rooms.GetGame(idRoom.Id)

	defer s.Sender.RoomInfo(game)

	return nil
}
