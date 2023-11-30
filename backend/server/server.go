package server

import (
	"encoding/json"
	"github.com/michelprogram/speeder-typer/rooms"
	"github.com/michelprogram/speeder-typer/utils"
	"golang.org/x/net/websocket"
	"io"
	"log"
)

type Server struct {
	Users  map[*websocket.Conn]string
	Events map[string]EventDataHandler
	Sender EventSender
	Rooms  *rooms.Rooms
}

func NewServer(sender EventSender, events map[string]EventDataHandler) *Server {

	return &Server{
		Users:  make(map[*websocket.Conn]string, 0),
		Rooms:  rooms.NewRooms(),
		Events: events,
		Sender: sender,
	}

}

func (s *Server) handleWebSocketEvents(data []byte, ws *websocket.Conn) error {

	var event EventJson

	err := json.Unmarshal(data, &event)

	if err != nil {
		return err
	}

	eventFunc, ok := s.Events[event.Name]

	if ok {
		err = eventFunc.HandleData(s, ws, event.Data)
		if err != nil {
			return err
		}
	} else {
		log.Printf("EventJson %s doesn't exist\n", event.Name)
	}

	return nil

}

func (s *Server) ReadLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Read error %s", err)
			continue
		}
		msg := buf[:n]
		go func() {
			err := s.handleWebSocketEvents(msg, ws)
			if err != nil {
				log.Println("Error during event " + err.Error())
			}
		}()
	}
}

func (s *Server) AddUser(ws *websocket.Conn) error {

	s.Users[ws] = utils.RandomUsername()

	err := s.Sender.RoomsInfo(s.Users, s.Rooms.Games)
	if err != nil {
		return err
	}

	return s.Sender.UserLogged(ws, s.Users[ws])

}

func (s *Server) RemoveUser(ws *websocket.Conn) error {

	username, exist := s.Users[ws]

	if exist {
		delete(s.Users, ws)
	}

	key := s.Rooms.FindUserAmongRooms(s.Users[ws])

	if key != "" {
		s.Rooms.LeaveRoom(username, key)
	}

	return s.Sender.RoomsInfo(s.Users, s.Rooms.Games)
}
