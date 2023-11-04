package types

type RoomInfo struct {
	Name    string  `json:"id"`
	Status  bool    `json:"status"`
	Players []*User `json:"users"`
}
