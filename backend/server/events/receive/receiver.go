package receive

import (
	"errors"
	"github.com/michelprogram/speeder-typer/server"
	"github.com/michelprogram/speeder-typer/types"
	"github.com/michelprogram/speeder-typer/utils"
	"golang.org/x/net/websocket"
	"strconv"
)

var (
	ROOM_START            = "room-start"
	ROOM_CREATED          = "room-created"
	ROOM_JOIN_BY_USERNAME = "room-join-by-username"
	PONG                  = "pong"
)

var _ server.EventReceiver = (*Receiver)(nil)

type Receiver struct {
	Sender server.EventSender
}

//TODO : change parameters to be a struct and jsonmarshal

func NewReceiver(sender server.EventSender) server.EventReceiver {
	return &Receiver{
		Sender: sender,
	}

}

func (r *Receiver) OnStartGame(websocketServer *server.Server, conn *websocket.Conn, parameters map[string]string) error {

	id, _ := parameters["id"]

	game := websocketServer.Rooms.Games[id]

	res := game.IsEveryoneReady()

	response := types.Response{
		Name: ROOM_START,
		Data: res,
	}

	err := r.Sender.JsonResponse(conn, response)

	if err != nil {
		return err
	}

	return nil
}

func (r *Receiver) OnSetReady(websocketServer *server.Server, conn *websocket.Conn, parameters map[string]string) error {

	id, _ := parameters["id"]
	username, _ := parameters["username"]
	ready, _ := parameters["ready"]

	readyBool, err := strconv.ParseBool(ready)

	if err != nil {
		return err
	}

	game := websocketServer.Rooms.Games[id]
	defer r.Sender.RoomInfo(conn, game)

	game.SetUserReady(username, readyBool)

	return nil
}

func (r *Receiver) OnPing(_ *server.Server, conn *websocket.Conn, _ map[string]string) error {
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

func (r *Receiver) OnCreateRoom(websocketServer *server.Server, conn *websocket.Conn, _ map[string]string) error {
	defer r.Sender.RoomsInfo(websocketServer.Users, websocketServer.Rooms.Games)

	var response types.Response

	id := utils.RandomIdRoom()

	user := types.NewUser(conn, websocketServer.Users[conn])

	err := websocketServer.Rooms.CreateRoom(user, id)
	if err != nil {
		return err
	}

	defer r.Sender.RoomInfo(conn, websocketServer.Rooms.Games[id])

	response = types.Response{
		Name: ROOM_CREATED,
		Data: id,
	}

	err = r.Sender.JsonResponse(conn, response)

	if err != nil {
		return err
	}

	return nil
}

func (r *Receiver) OnJoinRoom(websocketServer *server.Server, conn *websocket.Conn, parameters map[string]string) error {

	defer r.Sender.RoomsInfo(websocketServer.Users, websocketServer.Rooms.Games)

	var response types.Response

	id, _ := parameters["id"]

	username := websocketServer.Users[conn]

	if websocketServer.Rooms.IsUserInRoom(username, id) != nil {
		return nil
	}

	user := types.NewUser(conn, username)

	err := websocketServer.Rooms.JoinRoom(user, id)

	defer r.Sender.RoomInfo(conn, websocketServer.Rooms.Games[id])

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

	defer r.Sender.RoomsInfo(websocketServer.Users, websocketServer.Rooms.Games)
	id, exist := parameters["id"]

	if !exist {
		return errors.New("Id " + id + " doesn't exist")
	}

	username := websocketServer.Users[conn]

	websocketServer.Rooms.LeaveRoom(username, id)

	return nil
}

func (r *Receiver) OnJoinByUsername(websocketServer *server.Server, conn *websocket.Conn, parameters map[string]string) error {

	defer r.Sender.RoomsInfo(websocketServer.Users, websocketServer.Rooms.Games)

	response := types.Response{
		Name: ROOM_JOIN_BY_USERNAME,
	}

	username, exist := parameters["username"]

	if !exist {
		return errors.New("Username " + username + " doesn't exist")
	}

	id := websocketServer.Rooms.FindUserAmongRooms(username)

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
