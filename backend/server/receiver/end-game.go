package receiver

import (
	"encoding/json"
	"github.com/michelprogram/speeder-typer/rooms"
	"github.com/michelprogram/speeder-typer/server"
	"golang.org/x/net/websocket"
)

type EndGame struct{}

func (e EndGame) HandleData(s *server.Server, _ *websocket.Conn, data json.RawMessage) error {

	var idRoom JSONIdRoom

	err := json.Unmarshal(data, &idRoom)

	if err != nil {
		return err
	}

	game := s.Rooms.GetGame(idRoom.Id)

	s.Rooms.DeleteRoom(idRoom.Id)

	defer s.Sender.GameEnd(game)

	game.SetStatus(rooms.Finish)

	return nil
}
