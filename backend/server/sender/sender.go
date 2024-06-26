package sender

import (
	"encoding/json"
	"github.com/michelprogram/speeder-typer/rooms"
	"github.com/michelprogram/speeder-typer/server"
	"github.com/michelprogram/speeder-typer/types"
	"github.com/michelprogram/speeder-typer/utils"
	"golang.org/x/net/websocket"
)

var (
	GAME_CANT_START = "game-cant-start"
	GAME_END        = "game-end"

	STATS_APP = "stats-app"

	ROOMS_INFO            = "rooms-info"
	ROOM_INFO             = "room-info"
	ROOM_CREATED          = "room-created"
	ROOM_JOIN_BY_USERNAME = "room-join-by-username"

	TEXT_INFO = "text-info"

	USERNAME = "username"
	PONG     = "pong"
)

var _ server.EventSender = (*Sender)(nil)

type Sender struct{}

func NewSender() *Sender {
	return &Sender{}
}

func (s *Sender) JoinedRoomByUsername(conn *websocket.Conn, id string) error {

	response := types.Response{
		Name: ROOM_JOIN_BY_USERNAME,
	}

	if id == "" {
		response.Data = "-1"
	} else {
		response.Data = id
	}

	return s.JsonResponse(conn, response)

}

func (s *Sender) RoomCreated(conn *websocket.Conn, id string) error {
	response := types.Response{
		Name: ROOM_CREATED,
		Data: id,
	}

	return s.JsonResponse(conn, response)
}

func (s *Sender) Pong(conn *websocket.Conn) error {
	response := types.Response{
		Name: PONG,
		Data: "Pong",
	}

	return s.JsonResponse(conn, response)
}

func (s *Sender) GameCantStart(conn *websocket.Conn) error {
	response := types.Response{
		Name: GAME_CANT_START,
	}

	return s.JsonResponse(conn, response)

}

func (s *Sender) GameEnd(game *rooms.Game) error {

	response := types.Response{
		Name: GAME_END,
		Data: rooms.CreateScoresWithGame(game),
	}

	res, err := json.Marshal(response)

	if err != nil {
		return err
	}

	for _, user := range game.Users {

		_, err = user.Ws.Write(res)
		if err != nil {
			return err
		}
	}

	return nil

}

func (s *Sender) RoomInfo(game *rooms.Game) error {

	response := types.Response{
		Name: ROOM_INFO,
		Data: game,
	}

	res, err := json.Marshal(response)

	if err != nil {
		return err
	}

	for _, user := range game.Users {

		_, err = user.Ws.Write(res)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Sender) RoomsInfo(users map[*websocket.Conn]string, rooms map[string]*rooms.Game) error {

	stats := make([]*types.RoomInfo, 0)

	for key, game := range rooms {
		stats = append(stats, types.NewRoomInfo(key, game.Status, game.GetUsers()))
	}

	response := types.Response{
		Name: ROOMS_INFO,
		Data: NewStats(stats),
	}

	res, err := json.Marshal(response)

	if err != nil {
		return err
	}

	for user := range users {
		user.Write(res)
	}

	return nil
}

func (s *Sender) JsonResponse(ws *websocket.Conn, response types.Response) error {

	res, err := json.Marshal(response)

	if err != nil {
		return err
	}

	_, err = ws.Write(res)

	return err
}

func (s *Sender) UserLogged(ws *websocket.Conn, username string) error {

	response := types.Response{
		Name: USERNAME,
		Data: username,
	}

	return s.JsonResponse(ws, response)
}

func (s *Sender) CodeInfo(users []*types.User, textinfo server.TextInfo) error {
	response := types.Response{
		Name: TEXT_INFO,
		Data: textinfo,
	}

	res, err := json.Marshal(response)

	if err != nil {
		return err
	}

	for _, user := range users {

		_, err = user.Ws.Write(res)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Sender) StatsApp(ws *websocket.Conn, users map[*websocket.Conn]string) error {

	usernames := make([]string, len(users))

	i := 0

	for _, username := range users {
		usernames[i] = username
		i++
	}

	response := types.Response{
		Name: STATS_APP,
		Data: StatsApp{
			Players: usernames,
			Game:    utils.GAME_PLAYED,
		},
	}

	res, err := json.Marshal(response)

	if err != nil {
		return err
	}

	_, err = ws.Write(res)
	if err != nil {
		return err
	}

	return nil
}
