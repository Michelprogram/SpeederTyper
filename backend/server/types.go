package server

import (
	"github.com/michelprogram/speeder-typer/types"
	"golang.org/x/net/websocket"
)

type EventReceiver interface {
	OnPing(websocketServer *Server, conn *websocket.Conn, parameters map[string]string) error
	OnJoinRoom(websocketServer *Server, conn *websocket.Conn, parameters map[string]string) error
	OnLeaveRoom(websocketServer *Server, conn *websocket.Conn, parameters map[string]string) error
	OnJoinByUsername(websocketServer *Server, conn *websocket.Conn, parameters map[string]string) error
}

type EventSender interface {
	RoomsInfo(map[*websocket.Conn]string, map[string][]*types.User) error
	JsonResponse(ws *websocket.Conn, response types.Response) error
	UserLogged(ws *websocket.Conn, username string) error
}

type EventJson struct {
	Name string            `json:"name"`
	Data map[string]string `json:"data"`
}

type ReceiverFunction func(*Server, *websocket.Conn, map[string]string) error

type Handler struct {
	*Server
	*websocket.Conn
	Parameters map[string]string
}

type Server struct {
	Users  map[*websocket.Conn]string
	Events map[string]ReceiverFunction
	Sender EventSender
	*types.Room
}
