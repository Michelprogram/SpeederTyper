package receiver

import (
	"encoding/json"
	"github.com/michelprogram/speeder-typer/server"
	"golang.org/x/net/websocket"
)

type Typing struct{}

func (t Typing) HandleData(s *server.Server, _ *websocket.Conn, data json.RawMessage) error {

	var typerJSON JSONTyper

	err := json.Unmarshal(data, &typerJSON)

	if err != nil {
		return err
	}

	game := s.Rooms.GetGame(typerJSON.Id)
	defer s.Sender.RoomInfo(game)

	game.SetUserPosition(typerJSON.Username)

	return nil
}
