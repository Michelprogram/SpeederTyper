package rooms

import (
	"github.com/michelprogram/speeder-typer/types"
	"sort"
	"time"
	"unicode"
)

type Score struct {
	WordCompleted      int     `json:"wordCompleted"`
	LetterCompleted    int     `json:"letterCompleted"`
	CharacterPerSecond float64 `json:"characterPerSecond"`
	Classment          int     `json:"classment"`
	Username           string  `json:"username"`
}

func CreateScoresWithGame(game *Game) []*Score {

	scores := make([]*Score, len(game.Users))

	sort.Sort(types.ByPosition(game.Users))

	for index, user := range game.Users {
		score := &Score{
			WordCompleted:      computeWordsCompleted(game.Text[:user.Position]),
			LetterCompleted:    user.Position,
			CharacterPerSecond: computeCharacterPerSecond(user.Position, game.StartAt, game.EndAt),
			Classment:          index + 1,
			Username:           user.Username,
		}

		scores[index] = score
	}

	return scores

}

func computeWordsCompleted(text string) (counter int) {

	for _, char := range text {
		if unicode.IsSpace(char) {
			counter++
		}
	}

	return counter

}

func computeCharacterPerSecond(characterTyped int, start, end time.Time) float64 {
	return float64(characterTyped) / end.Sub(start).Seconds()
}
