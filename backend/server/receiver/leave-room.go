package receiver

import (
	"encoding/json"
	"fmt"
	"github.com/michelprogram/speeder-typer/server"
	"golang.org/x/net/websocket"
)

type LeaveRoom struct{}

func (l LeaveRoom) HandleData(s *server.Server, ws *websocket.Conn, data json.RawMessage) error {

	var idRoom JSONIdRoom

	err := json.Unmarshal(data, &idRoom)

	if err != nil {
		return err
	}

	defer s.Sender.RoomsInfo(s.Users, s.Rooms.Games)

	username := s.Users[ws]

	s.Rooms.LeaveRoom(username, idRoom.Id)

	fmt.Printf("%v\n", s.Rooms.GetGame(idRoom.Id))

	return nil
}
