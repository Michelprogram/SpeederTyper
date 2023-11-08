package utils

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/google/uuid"
)

var Env string
var counter int
var Port int

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

func FetchRandomText() (sentence string, _ error) {
	response, err := http.Get("https://randomwordgenerator.com/json/words_ws.json")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var randomWordGenerator RandomWordGenerator
	err = json.Unmarshal(body, &randomWordGenerator)
	if err != nil {
		return "", err
	}

	rand.Seed(time.Now().UnixNano())

	min := 0
	max := len(randomWordGenerator.Data)

	for i := 0; i < 2; i++ {
		index := rand.Intn(max-min+1) + min

		sentence = fmt.Sprintf("%s %s", sentence, randomWordGenerator.Data[index].Word.Value)
	}

	//Remove first space
	return sentence[1:], nil
}
