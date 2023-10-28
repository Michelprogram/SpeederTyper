package handlers

import (
	"github.com/michelprogram/speeder-typer/server"
	"golang.org/x/net/websocket"
	"log"
)

func MainWS(ws *websocket.Conn, websocketServer *server.Server) error {

	log.Printf("New connection from client %s\n", ws.RemoteAddr())

	websocketServer.AddUser(ws)

	err := websocketServer.Sender.RoomsInfo(websocketServer.Users, websocketServer.Rooms)
	if err != nil {
		return err
	}

	websocketServer.ReadLoop(ws)

	log.Printf("User %s deconnected\n", ws.RemoteAddr())

	websocketServer.RemoveUser(ws)

	return nil
}
