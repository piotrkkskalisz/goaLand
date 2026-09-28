package transform

import "backend/internal/state/table"

type ExtendedClub struct {
	Position      int                 `json:"position"`
	TeamID        int                 `json:"id"`
	TeamName      string              `json:"name"`
	TeamCode      string              `json:"code"`
	CompetitionID int                 `json:"competitionID"`
	Points        int                 `json:"points"`
	Wins          int                 `json:"wins"`
	Draws         int                 `json:"draws"`
	Losses        int                 `json:"losses"`
	GoalsScored   int                 `json:"goalsScored"`
	GoalsConceded int                 `json:"goalsConceded"`
	Form          []table.MatchResult `json:"form"`
	IsLive        bool                `json:"isLive"`
}

func createExtendedClub(club *table.Club, position, competitionID int) ExtendedClub {
	return ExtendedClub{
		Position:      position,
		TeamID:        club.TeamID,
		TeamName:      club.TeamName,
		TeamCode:      club.TeamCode,
		CompetitionID: competitionID,
		Points:        club.Points,
		Wins:          club.Wins,
		Draws:         club.Draws,
		Losses:        club.Losses,
		GoalsScored:   club.GoalsScored,
		GoalsConceded: club.GoalsConceded,
		Form:          club.Form,
		IsLive:        club.IsLive,
	}
}

func GetFullInformationClubs(clubs []*table.Club, competitionID int) []ExtendedClub {
	extendedClubs := make([]ExtendedClub, 0, len(clubs))
	for index, club := range clubs {
		extendedClubs = append(
			extendedClubs,
			createExtendedClub(club, index+1, competitionID),
		)
	}
	return extendedClubs
}

func GetMiniTableForClub(
	clubs []*table.Club, competitionID int, clubID int,
) []ExtendedClub {
	extendedClubs := make([]ExtendedClub, 0, len(clubs))

	for index, club := range clubs {
		extendedClubs = append(
			extendedClubs,
			createExtendedClub(club, index+1, competitionID),
		)
	}
	return extendedClubs
}
