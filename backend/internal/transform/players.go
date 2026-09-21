package transform

import (
	"backend/internal/database"
	"backend/internal/utils"
)

type playersOnPosition struct {
	Position string   `json:"position"`
	Players  []Player `json:"players"`
}

type Player struct {
	PlayerID    int    `json:"playerID"`
	PlayerName  string `json:"playerName"`
	Nationality string `json:"nationality"`
}

func GroupByPositions(seasonPlayers []database.SeasonPlayer) []playersOnPosition {
	groupPlayers := make(map[string][]Player)
	for _, position := range utils.Positions {
		groupPlayers[position] = make([]Player, 0)
	}
	for _, seasonPlayer := range seasonPlayers {
		player := seasonPlayer.Player
		if _, ok := groupPlayers[player.Position]; ok {
			groupPlayers[player.Position] = append(groupPlayers[player.Position], Player{
				PlayerID:    player.PlayerID,
				PlayerName:  player.Name,
				Nationality: player.NationalityArea.Name,
			})
		}
	}

	var playersResponse []playersOnPosition

	for _, position := range utils.Positions {
		playersResponse = append(playersResponse, playersOnPosition{
			Players:  groupPlayers[position],
			Position: position,
		})
	}

	return playersResponse
}
