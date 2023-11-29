package receiver

import (
	"encoding/json"
	"github.com/michelprogram/speeder-typer/server"
	"golang.org/x/net/websocket"
)

type StatsApp struct{}

func (c StatsApp) HandleData(s *server.Server, ws *websocket.Conn, _ json.RawMessage) error {
	return s.Sender.StatsApp(ws, s.Users)
}
