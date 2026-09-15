package transform

import (
	"backend/internal/database"
	"backend/internal/utils"
	"cmp"
	"slices"
)

type CompetitionEdition struct {
	CompetitionID int `json:"competitionId"`

	Name            string `json:"name"`
	Code            string `json:"code"`
	CompetitionType string `json:"competitionType"`

	StartYear int    `json:"startYear"`
	Status    string `json:"status"`
}

type GoalScorerResponse struct {
	PlayerID   int    `json:"playerID"`
	PlayerName string `json:"playerName"`
	TeamName   string `json:"teamName"`

	Position int `json:"position"`
	Goals    int `json:"goals"`
	Assists  int `json:"assists"`
}

func createCompetitionEdition(competition database.Competition, edition database.Edition) CompetitionEdition {
	return CompetitionEdition{
		CompetitionID: competition.CompetitionID,

		Name:            competition.Name,
		Code:            competition.Code,
		CompetitionType: competition.CompetitionType,

		StartYear: edition.StartYear,
		Status:    edition.Status,
	}
}

func getAcitveOrUpcoming(competition database.Competition) (CompetitionEdition, bool) {
	for _, edition := range competition.Editions {
		if utils.IsCurrent(edition.Status) {
			return createCompetitionEdition(competition, edition), true
		}
	}

	return CompetitionEdition{}, false
}

func GetCompetitionEdition(competitions []database.Competition) []CompetitionEdition {
	var response []CompetitionEdition

	for _, competition := range competitions {
		if competitionEdition, ok := getAcitveOrUpcoming(competition); ok {
			response = append(response, competitionEdition)
		}
	}

	return response
}

func createGoalScorerResponse(player database.SeasonPlayer, position int) GoalScorerResponse {
	return GoalScorerResponse{
		PlayerID:   player.PlayerID,
		PlayerName: player.Player.Name,
		TeamName:   player.Team.FullName,
		Position:   position,
		Goals:      *player.Goals,
		Assists:    utils.IntOrZero(player.Assists),
	}
}
func GetEditionGoalScorers(players []database.SeasonPlayer) []GoalScorerResponse {
	if len(players) == 0 {
		return nil
	}

	slices.SortFunc(players, func(a, b database.SeasonPlayer) int {
		if result := cmp.Compare(*b.Goals, *a.Goals); result != 0 {
			return result
		}
		return cmp.Compare(a.PlayerID, b.PlayerID)
	})

	var goalScorers []GoalScorerResponse
	position := 1

	goalScorers = append(goalScorers, createGoalScorerResponse(players[0], position))

	for i := range len(players) - 1 {
		if *players[i].Goals != *players[i+1].Goals {
			position += 1
		}
		goalScorers = append(goalScorers, createGoalScorerResponse(players[i+1], position))

	}

	return goalScorers

}
