package main

import (
	"github.com/michelprogram/speeder-typer/handlers"
	"github.com/michelprogram/speeder-typer/server"
	"github.com/michelprogram/speeder-typer/server/events/receive"
	"github.com/michelprogram/speeder-typer/server/events/send"
	"github.com/michelprogram/speeder-typer/utils"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {

	log.Printf("Environnement %s\n", utils.Env)

	log.Println("🚀 Lancement de l'api sur le port 3000...")

	sender := send.NewSender()
	receiver := receive.NewReceiver(sender)
	websocketServer := server.NewServer(sender, receiver)

	http.Handle("/ws", websocket.Handler(func(ws *websocket.Conn) {
		err := handlers.MainWS(ws, websocketServer)
		if err != nil {
			log.Printf(err.Error())
		}
	}))

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err.Error())
	}
}
