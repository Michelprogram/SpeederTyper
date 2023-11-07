package types

type RoomInfo struct {
	Name    string  `json:"id"`
	Status  int     `json:"status"`
	Players []*User `json:"users"`
}

func NewRoomInfo(name string, status int, players []*User) *RoomInfo {
	return &RoomInfo{
		name, status, players,
	}
}
