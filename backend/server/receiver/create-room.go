package receiver

import (
	"encoding/json"
	"github.com/michelprogram/speeder-typer/server"
	"github.com/michelprogram/speeder-typer/types"
	"github.com/michelprogram/speeder-typer/utils"
	"golang.org/x/net/websocket"
)

type CreateRoom struct{}

func (c CreateRoom) HandleData(s *server.Server, ws *websocket.Conn, _ json.RawMessage) error {

	defer s.Sender.RoomsInfo(s.Users, s.Rooms.Games)

	id := utils.RandomIdRoom()

	user := types.NewUser(ws, s.Users[ws])

	err := s.Rooms.CreateRoom(user, id)
	if err != nil {
		return err
	}

	defer s.Sender.RoomInfo(s.Rooms.GetGame(id))

	return s.Sender.RoomCreated(ws, id)
}
