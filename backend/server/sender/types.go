package sender

import "github.com/michelprogram/speeder-typer/types"

type Stats struct {
	Stats []*types.RoomInfo `json:"stats"`
}

func NewStats(stats []*types.RoomInfo) *Stats {
	return &Stats{
		stats,
	}
}

type StatsApp struct {
	Players []string `json:"players"`
	Game    int      `json:"game"`
}
