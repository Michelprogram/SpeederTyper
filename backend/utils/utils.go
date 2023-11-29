package utils

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

var Env string
var counter int
var Port int
var GAME_PLAYED = 0

func init() {

	flag.StringVar(&Env, "env", "prod", "Help for debugging")
	flag.IntVar(&Port, "port", 3000, "Port the application expose")
	flag.Parse()
}

func RandomIdRoom() string {
	if Env == "dev" {
		return "room-id"
	}

	return uuid.NewString()
}

func RandomUsername() string {

	if Env == "dev" {
		counter++
		return fmt.Sprintf("username - %d", counter)
	}

	rand.Seed(time.Now().UnixNano())
	min := 0

	username := fmt.Sprintf("%s %s", PROGRAMMING_NAME[rand.Intn(SIZE-min+1)+min], ADJECTIVES[rand.Intn(SIZE_2-min+1)+min])

	return username
}
