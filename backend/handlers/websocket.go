package handlers

import (
	"github.com/michelprogram/speeder-typer/server"
	"golang.org/x/net/websocket"
	"log"
)

func MainWS(ws *websocket.Conn, websocketServer *server.Server) error {

	log.Printf("New connection from client %s\n", ws.RemoteAddr())

	err := websocketServer.AddUser(ws)
	if err != nil {
		return err
	}

	websocketServer.ReadLoop(ws)

	return websocketServer.RemoveUser(ws)
}
