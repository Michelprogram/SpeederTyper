package pkg

import (
	"encoding/json"
	"errors"
	"github.com/michelprogram/speeder-typer/utils"
	"golang.org/x/net/websocket"
)

var (
	ROOM_INFO             = "room-info"
	ROOM_CREATED          = "room-created"
	ROOM_JOIN_BY_USERNAME = "room-join-by-username"
	USERNAME              = "username"
	PONG                  = "pong"
)

type RoomInfo struct {
	Name    string  `json:"id"`
	Players []*User `json:"players"`
}

type Event struct {
	Name string            `json:"name"`
	Data map[string]string `json:"data"`
}

type Response struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

func SendJsonResponse(ws *websocket.Conn, response Response) error {

	res, err := json.Marshal(response)

	if err != nil {
		return err
	}

	ws.Write(res)

	return nil
}

func OnPing(server *Server, user *User, params map[string]string) error {
	response := Response{
		Name: PONG,
		Data: "Pong",
	}

	err := SendJsonResponse(user.Ws, response)

	if err != nil {
		return err
	}
	return nil
}

func OnJoinRoom(server *Server, user *User, params map[string]string) error {
	defer SendRoomsInfo(server, user, params)

	var response Response

	id, exist := params["id"]
	if !exist {
		id = utils.RandomIdRoom()

		response = Response{
			Name: ROOM_CREATED,
			Data: id,
		}
	}

	if server.Room.IsUserInRoom(user.Username, id) != nil {
		return nil
	}

	err := server.Room.JoinRoom(user, id)

	if err != nil {
		return err
	}

	err = SendJsonResponse(user.Ws, response)

	if err != nil {
		return err
	}

	return nil
}

func OnLeaveRoom(server *Server, user *User, params map[string]string) error {

	defer SendRoomsInfo(server, user, params)
	id, exist := params["id"]

	if !exist {
		return errors.New("Id " + id + " doesn't exist")
	}

	server.Room.LeaveRoom(user, id)

	return nil
}

func OnJoinByUsername(server *Server, user *User, params map[string]string) error {

	defer SendRoomsInfo(server, user, params)

	response := Response{
		Name: ROOM_JOIN_BY_USERNAME,
	}

	username, exist := params["username"]

	if !exist {
		return errors.New("Username " + username + " doesn't exist")
	}

	id := server.Room.FindUserAmongRooms(username)

	if id == "" {
		response.Data = "-1"
	} else {
		response.Data = id
	}

	err := SendJsonResponse(user.Ws, response)

	if err != nil {
		return err
	}

	return nil
}

func SendRoomsInfo(server *Server, _ *User, params map[string]string) error {

	stats := make([]RoomInfo, 0)

	for key, value := range server.Rooms {
		stats = append(stats, RoomInfo{
			Name:    key,
			Players: value,
		})
	}

	response := Response{
		Name: ROOM_INFO,
		Data: struct {
			Stats []RoomInfo `json:"stats"`
			Users []string   `json:"users"`
		}{stats, Users},
	}

	res, err := json.Marshal(response)

	if err != nil {
		return err
	}

	for user := range server.Conns {
		user.Write(res)
	}

	return nil

}
