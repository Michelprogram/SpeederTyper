package main

import (
	"fmt"
	"github.com/michelprogram/speeder-typer/handlers"
	"github.com/michelprogram/speeder-typer/server"
	"github.com/michelprogram/speeder-typer/server/receiver"
	"github.com/michelprogram/speeder-typer/server/sender"
	"github.com/michelprogram/speeder-typer/utils"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {

	log.Printf("Environnement %s\n", utils.Env)

	log.Printf("🚀 Lancement de l'api sur le port %d\n", utils.Port)

	sender := sender.NewSender()

	events := map[string]server.EventDataHandler{
		"ping":             &receiver.PingHandler{},
		"create-room":      &receiver.CreateRoom{},
		"join-room":        &receiver.JoinRoom{},
		"leave-room":       &receiver.LeaveRoom{},
		"join-by-username": &receiver.JoinByUsername{},
		"set-ready":        &receiver.SetReady{},
		"start-game":       &receiver.StartGame{},
		"typing-game":      &receiver.Typing{},
		"end-game":         &receiver.EndGame{},
	}

	websocketServer := server.NewServer(sender, events)

	http.Handle("/ws", websocket.Handler(func(ws *websocket.Conn) {
		err := handlers.MainWS(ws, websocketServer)
		if err != nil {
			log.Printf(err.Error())
		}
	}))

	expose := fmt.Sprintf(":%d", utils.Port)

	err := http.ListenAndServe(expose, nil)
	if err != nil {
		panic(err.Error())
	}
}
