package receiver

import (
	"encoding/json"
	"github.com/michelprogram/speeder-typer/server"
	"golang.org/x/net/websocket"
)

type JoinByUsername struct{}

func (j JoinByUsername) HandleData(s *server.Server, ws *websocket.Conn, data json.RawMessage) error {

	var usernameJson JSONUsername

	err := json.Unmarshal(data, &usernameJson)

	if err != nil {
		return err
	}

	defer s.Sender.RoomsInfo(s.Users, s.Rooms.Games)

	id := s.Rooms.FindUserAmongRooms(usernameJson.Username)

	return s.Sender.JoinedRoomByUsername(ws, id)
}
