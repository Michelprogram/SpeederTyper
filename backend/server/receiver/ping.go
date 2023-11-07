package receiver

import (
	"encoding/json"
	"github.com/michelprogram/speeder-typer/server"
	"golang.org/x/net/websocket"
)

type PingHandler struct{}

func (p *PingHandler) HandleData(s *server.Server, ws *websocket.Conn, _ json.RawMessage) error {

	return s.Sender.Pong(ws)

}
