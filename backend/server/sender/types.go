package sender

import "github.com/michelprogram/speeder-typer/types"

type Stats struct {
	Stats []*types.RoomInfo `json:"stats"`
	Users []string          `json:"users"`
}

func NewStats(stats []*types.RoomInfo, usernames []string) *Stats {
	return &Stats{
		stats, usernames,
	}
}
