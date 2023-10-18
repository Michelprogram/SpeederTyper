package utils

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func RandomIdRoom() string {
	return uuid.NewString()
}

func RandomUsername() string {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := SIZE

	username := ""

	for i := 0; i < 1; i++ {
		index := rand.Intn(max-min+1) + min
		username += USERNAMES[index]
	}

	return username
}

func FindAndRemoveElement(array []string, looking string) []string {
	for index, el := range array {
		if el == looking {
			array[index] = array[len(array)-1]
			return array[:len(array)-1]
		}
	}
	return array
}
