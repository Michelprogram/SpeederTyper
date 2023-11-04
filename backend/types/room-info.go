package types

type RoomInfo struct {
	Name    string  `json:"id"`
	Players []*User `json:"users"`
}
