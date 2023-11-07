package server

import (
	"encoding/json"
	"github.com/michelprogram/speeder-typer/rooms"
	"github.com/michelprogram/speeder-typer/types"
	"golang.org/x/net/websocket"
)

type EventSender interface {
	RoomsInfo(map[*websocket.Conn]string, map[string]*rooms.Game) error
	GameEnd(game *rooms.Game) error
	GameCantStart(conn *websocket.Conn) error
	RoomInfo(*rooms.Game) error
	Pong(conn *websocket.Conn) error
	RoomCreated(conn *websocket.Conn, id string) error
	JoinedRoomByUsername(conn *websocket.Conn, id string) error
	JsonResponse(ws *websocket.Conn, response types.Response) error
	UserLogged(ws *websocket.Conn, username string) error
}

type EventJson struct {
	Name string          `json:"name"`
	Data json.RawMessage `json:"data"`
}

type EventDataHandler interface {
	HandleData(s *Server, ws *websocket.Conn, data json.RawMessage) error
}
