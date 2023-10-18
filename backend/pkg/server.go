package pkg

import (
	"encoding/json"
	"io"
	"log"

	"github.com/michelprogram/speeder-typer/utils"
	"golang.org/x/net/websocket"
)

type Server struct {
	Conns map[*websocket.Conn]string
	*Room
	Events map[string]EventHandler
}

type EventHandler func(*Server, *User, map[string]string) error

var Users []string

func NewServer() *Server {

	events := make(map[string]EventHandler)

	//Essayer avec les reflects
	events["ping"] = OnPing
	events["join-room"] = OnJoinRoom
	events["leave-room"] = OnLeaveRoom
	events["join-by-username"] = OnJoinByUsername

	Users = make([]string, 0)

	return &Server{
		Conns:  make(map[*websocket.Conn]string),
		Room:   NewRooms(),
		Events: events,
	}

}

func (s *Server) TriggerEvent(data []byte, ws *websocket.Conn) error {

	var event Event

	err := json.Unmarshal(data, &event)

	if err != nil {
		return err
	}

	eventFunc, ok := s.Events[event.Name]

	user := &User{
		Ws:       ws,
		Username: s.Conns[ws],
	}

	if ok {
		err = eventFunc(s, user, event.Data)
		if err != nil {
			return err
		}
	} else {
		log.Printf("Event %s doesn't exist\n", event.Name)
	}

	return nil

}

func (s *Server) MainWS(ws *websocket.Conn) {
	log.Printf("New connection from client %s\n", ws.RemoteAddr())

	s.Conns[ws] = utils.RandomUsername()

	user := NewUser(ws, s.Conns[ws])

	Users = append(Users, user.Username)

	SendRoomsInfo(s, user, map[string]string{})

	s.readLoop(ws)

	log.Printf("User %s deconnected\n", ws.RemoteAddr())

	Users = utils.FindAndRemoveElement(Users, user.Username)
}

func (s *Server) readLoop(ws *websocket.Conn) {
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
			err := s.TriggerEvent(msg, ws)
			if err != nil {
				log.Println("Error during event " + err.Error())
			}
		}()
	}
}
