package rooms

import (
	"encoding/json"
	"github.com/michelprogram/speeder-typer/types"
	"sync"
)

type Rooms struct {
	sync.RWMutex
	Games map[string]*Game
}

func NewRooms() *Rooms {
	return &Rooms{
		Games: make(map[string]*Game),
	}
}

func (r *Rooms) BroadcastToRoom(response types.Response, key string) error {

	res, err := json.Marshal(response)

	if err != nil {
		return err
	}

	r.RLock()
	game := r.Games[key]
	r.Unlock()

	for _, user := range game.Users {
		_, err = user.Ws.Write(res)
		if err != nil {
			return err
		}
	}

	return nil

}

func (r *Rooms) CreateRoom(user *types.User, key string) error {
	r.RLock()
	defer r.RUnlock()

	game, err := NewGame(key, user.Username)

	if err != nil {
		return err
	}

	game.UserJoin(user)

	r.Games[key] = game

	return nil
}

func (r *Rooms) JoinRoom(user *types.User, key string) error {
	r.RLock()
	defer r.RUnlock()

	game := r.Games[key]

	game.UserJoin(user)

	return nil
}

func (r *Rooms) LeaveRoom(username string, key string) error {
	r.RLock()
	defer r.RUnlock()

	game := r.Games[key]

	size := game.UserLeft(username)

	if size == 0 {
		r.DeleteRoom(key)
	}

	return nil
}

func (r *Rooms) DeleteRoom(key string) {
	r.RLock()
	defer r.RUnlock()

	_, ok := r.Games[key]
	if ok {
		delete(r.Games, key)
	}
}

func (r *Rooms) GetUsernames(key string) []string {

	r.RLock()
	game := r.Games[key]
	r.RUnlock()

	return game.Usernames()
}

func (r *Rooms) GetUsers(key string) []*types.User {

	r.RLock()
	defer r.RUnlock()

	game, ok := r.Games[key]
	if ok {
		return game.GetUsers()
	}

	return nil

}

func (r *Rooms) IsUserInRoom(username string, key string) *types.User {
	r.RLock()
	defer r.RUnlock()

	game := r.Games[key]

	return game.IsUserInGame(username)
}

func (r *Rooms) FindUserAmongRooms(username string) string {
	r.RLock()
	defer r.RUnlock()

	for key, game := range r.Games {
		if game.IsUserInGame(username) != nil {
			return key
		}
	}

	return ""
}
