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

type GoalScorer struct {
	Position int `json:"position"`
	Goals    int `json:"goals"`
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

func GetEditionGoalScorers(players []database.SeasonPlayer) []GoalScorer {
	if len(players) == 0 {
		return nil
	}

	slices.SortFunc(players, func(a, b database.SeasonPlayer) int {
		if result := cmp.Compare(*b.Goals, *a.Goals); result != 0 {
			return result
		}
		return cmp.Compare(a.PlayerID, b.PlayerID)
	})

	var goalScorers []GoalScorer
	position := 1

	goalScorers = append(goalScorers, GoalScorer{
		Position: position,
		Goals:    *players[0].Goals,
	})

	for i := range len(players) - 1 {
		if *players[i].Goals != *players[i+1].Goals {
			position += 1
		}
		goalScorers = append(goalScorers, GoalScorer{
			Position: position,
			Goals:    *players[i+1].Goals,
		})
	}

	return goalScorers

}
