package pkg

import (
	"encoding/json"
	"sync"
)

type Room struct {
	sync.RWMutex
	Rooms map[string][]*User
}

func NewRooms() *Room {
	return &Room{
		Rooms: make(map[string][]*User),
	}
}

func (r *Room) BroadcastToRoom(response Response, key string) error {

	res, err := json.Marshal(response)

	if err != nil {
		return err
	}

	r.RLock()
	users := r.Rooms[key]
	r.Unlock()

	for _, user := range users {
		user.Ws.Write(res)
	}

	return nil

}

func (r *Room) JoinRoom(user *User, key string) error {
	r.RLock()
	defer r.RUnlock()

	r.Rooms[key] = append(r.Rooms[key], user)

	return nil
}

func (r *Room) LeaveRoom(currentUser *User, key string) error {
	r.RLock()
	defer r.RUnlock()

	users, _ := r.Rooms[key]

	for index, user := range users {
		if user.Username == currentUser.Username {
			r.Rooms[key] = append(r.Rooms[key][:index], r.Rooms[key][index+1:]...)
			break
		}
	}

	if len(r.Rooms[key]) == 0 {
		r.DeleteRoom(key)
	}

	return nil
}

func (r *Room) DeleteRoom(key string) {
	r.RLock()
	defer r.RUnlock()

	_, ok := r.Rooms[key]
	if ok {
		delete(r.Rooms, key)
	}
}

func (r Room) GetUsernames(key string) []string {

	var usernames []string

	r.RLock()
	users := r.Rooms[key]
	r.RUnlock()

	for _, user := range users {
		usernames = append(usernames, user.Username)
	}

	return usernames
}

func (r Room) GetUsers(key string) []*User {

	r.RLock()
	defer r.RUnlock()

	user, ok := r.Rooms[key]
	if ok {
		return user
	}

	return nil

}

func (r Room) IsUserInRoom(username string, key string) *User {
	r.RLock()
	defer r.RUnlock()

	users := r.GetUsers(key)

	for _, user := range users {
		if user.Username == username {
			return user
		}
	}

	return nil
}

func (r Room) FindUserAmongRooms(user string) string {
	r.RLock()
	defer r.RUnlock()

	for key, _ := range r.Rooms {
		if r.IsUserInRoom(user, key) != nil {
			return key
		}
	}

	return ""
}
