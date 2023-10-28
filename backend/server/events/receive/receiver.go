package receive

import (
	"errors"
	"github.com/michelprogram/speeder-typer/server"
	"github.com/michelprogram/speeder-typer/types"
	"github.com/michelprogram/speeder-typer/utils"
	"golang.org/x/net/websocket"
)

var (
	ROOM_CREATED          = "room-created"
	ROOM_JOIN_BY_USERNAME = "room-join-by-username"
	PONG                  = "pong"
)

var _ server.EventReceiver = (*Receiver)(nil)

type Receiver struct {
	Sender server.EventSender
}

func NewReceiver(sender server.EventSender) server.EventReceiver {
	return &Receiver{
		Sender: sender,
	}
}

func (r *Receiver) OnPing(webSocketServer *server.Server, conn *websocket.Conn, parameters map[string]string) error {
	response := types.Response{
		Name: PONG,
		Data: "Pong",
	}

	err := r.Sender.JsonResponse(conn, response)

	if err != nil {
		return err
	}
	return nil
}

func (r *Receiver) OnJoinRoom(websocketServer *server.Server, conn *websocket.Conn, parameters map[string]string) error {

	defer r.Sender.RoomsInfo(websocketServer.Users, websocketServer.Rooms)

	var response types.Response

	id, exist := parameters["id"]
	if !exist {
		id = utils.RandomIdRoom()

		response = types.Response{
			Name: ROOM_CREATED,
			Data: id,
		}
	}

	username := websocketServer.Users[conn]

	if websocketServer.IsUserInRoom(username, id) != nil {
		return nil
	}

	user := types.NewUser(conn, username)

	err := websocketServer.JoinRoom(user, id)

	if err != nil {
		return err
	}

	err = r.Sender.JsonResponse(conn, response)

	if err != nil {
		return err
	}

	return nil
}

func (r *Receiver) OnLeaveRoom(websocketServer *server.Server, conn *websocket.Conn, parameters map[string]string) error {

	defer r.Sender.RoomsInfo(websocketServer.Users, websocketServer.Rooms)
	id, exist := parameters["id"]

	if !exist {
		return errors.New("Id " + id + " doesn't exist")
	}

	username := websocketServer.Users[conn]

	websocketServer.Room.LeaveRoom(username, id)

	return nil
}

func (r *Receiver) OnJoinByUsername(websocketServer *server.Server, conn *websocket.Conn, parameters map[string]string) error {

	defer r.Sender.RoomsInfo(websocketServer.Users, websocketServer.Rooms)

	response := types.Response{
		Name: ROOM_JOIN_BY_USERNAME,
	}

	username, exist := parameters["username"]

	if !exist {
		return errors.New("Username " + username + " doesn't exist")
	}

	id := websocketServer.FindUserAmongRooms(username)

	if id == "" {
		response.Data = "-1"
	} else {
		response.Data = id
	}

	err := r.Sender.JsonResponse(conn, response)

	if err != nil {
		return err
	}

	return nil
}
