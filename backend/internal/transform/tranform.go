package transform

import (
	"backend/internal/database"
	"backend/internal/utils"
)

type CompetitionEdition struct {
	CompetitionID int `json:"competitionId"`

	Name            string `json:"name"`
	Code            string `json:"code"`
	CompetitionType string `json:"competitionType"`

	StartYear int    `json:"startYear"`
	Status    string `json:"status"`
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
