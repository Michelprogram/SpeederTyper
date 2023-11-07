package receiver

import (
	"encoding/json"
	"github.com/michelprogram/speeder-typer/server"
	"golang.org/x/net/websocket"
)

type SetReady struct{}

func (sr SetReady) HandleData(s *server.Server, _ *websocket.Conn, data json.RawMessage) error {

	var setReadyJson JSONSetReady

	err := json.Unmarshal(data, &setReadyJson)

	if err != nil {
		return err
	}

	game := s.Rooms.Games[setReadyJson.Id]
	defer s.Sender.RoomInfo(game)

	game.SetUserReady(setReadyJson.Username, setReadyJson.Ready)

	return nil
}
