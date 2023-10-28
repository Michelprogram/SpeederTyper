package server

import (
	"encoding/json"
	"github.com/michelprogram/speeder-typer/types"
	"github.com/michelprogram/speeder-typer/utils"
	"golang.org/x/net/websocket"
	"io"
	"log"
)

func NewServer(sender EventSender, receiver EventReceiver) *Server {

	eventsToReceive := map[string]ReceiverFunction{
		"ping":             receiver.OnPing,
		"join-room":        receiver.OnJoinRoom,
		"leave-room":       receiver.OnLeaveRoom,
		"join-by-username": receiver.OnJoinByUsername,
	}

	return &Server{
		Users:  make(map[*websocket.Conn]string),
		Room:   types.NewRooms(),
		Events: eventsToReceive,
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
		err = eventFunc(s, ws, event.Data)
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

func (s *Server) AddUser(ws *websocket.Conn) {

	s.Users[ws] = utils.RandomUsername()

	s.Sender.UserLogged(ws, s.Users[ws])
}

func (s *Server) RemoveUser(ws *websocket.Conn) {

	_, exist := s.Users[ws]

	if exist {
		delete(s.Users, ws)
	}

}
