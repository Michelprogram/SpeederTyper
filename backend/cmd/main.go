package main

import (
	"log"
	"net/http"

	"github.com/michelprogram/speeder-typer/pkg"
	"golang.org/x/net/websocket"
)

func main() {

	server := pkg.NewServer()

	log.Println("🚀 Lancement de l'api sur le port 3000...")

	http.Handle("/ws", websocket.Handler(server.MainWS))
	http.ListenAndServe(":3000", nil)
}
