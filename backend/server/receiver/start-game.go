package receiver

import (
	"encoding/json"
	"github.com/michelprogram/speeder-typer/github"
	"github.com/michelprogram/speeder-typer/server"
	"golang.org/x/net/websocket"
)

type StartGame struct{}

func (sg StartGame) HandleData(s *server.Server, ws *websocket.Conn, data json.RawMessage) error {

	var idRoom JSONIdRoom

	err := json.Unmarshal(data, &idRoom)

	if err != nil {
		return err
	}

	canStartGame := s.Rooms.IsEveryoneReady(idRoom.Id)

	//Si pas ready on informe que le creator
	if !canStartGame {
		return s.Sender.GameCantStart(ws)
	}

	//Si ready on informe tout le monde
	game := s.Rooms.GetGame(idRoom.Id)

	text, err := github.FetchRandomCode(game.Users, s.Sender)
	if err != nil {
		return err
	}

	game.SetText(text)

	defer s.Sender.RoomInfo(game)
	defer s.Sender.RoomsInfo(s.Users, s.Rooms.Games)

	s.Rooms.StartGame(idRoom.Id)

	return nil
}
