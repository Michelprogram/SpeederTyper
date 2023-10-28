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
