package transform

import (
	"backend/internal/database"
	"backend/internal/utils"
	"slices"
	"time"
)

type MatchResponse struct {
	MatchID int `json:"id"`

	StartTime time.Time `json:"startTime"`
	Status    string    `json:"status"`

	HomeTeam string `json:"homeTeam"`
	AwayTeam string `json:"awayTeam"`

	HomeScore *int `json:"homeScore,omitempty"`
	AwayScore *int `json:"awayScore,omitempty"`
}

type RoundResponse struct {
	Stage    string `json:"stage"`
	Matchday *int   `json:"matchday"`
}

type RoundMatchesResponse struct {
	Round   RoundResponse   `json:"round"`
	Matches []MatchResponse `json:"matches"`
}

type roundKey struct {
	stage    string
	matchday int
}

func matchStatus(status string) string {
	if slices.Contains(utils.UpcomingMatchStatuses, status) {
		return "scheduled"
	} else if slices.Contains(utils.LiveMatchStatuses, status) {
		return "live"
	} else if slices.Contains(utils.FinishedMatchStatuses, status) {
		return "finished"
	} else {
		return status
	}
}

func createMatchResponse(match database.Match) MatchResponse {
	return MatchResponse{
		MatchID: match.MatchID,

		StartTime: match.StartTime,
		Status:    matchStatus(match.Status),

		HomeTeam: match.HomeTeam.FullName,
		AwayTeam: match.AwayTeam.FullName,

		HomeScore: match.HomeGoals,
		AwayScore: match.AwayGoals,
	}
}

func GroupByRounds(matches []database.Match) []RoundMatchesResponse {
	groupedMatches := make([]RoundMatchesResponse, 0)
	var lastRound roundKey

	for _, match := range matches {
		matchday := 0
		if match.Matchday != nil {
			matchday = *match.Matchday
		}

		round := roundKey{
			stage:    match.Stage,
			matchday: matchday,
		}
		response := createMatchResponse(match)

		if len(groupedMatches) == 0 || lastRound != round {
			groupedMatches = append(groupedMatches, RoundMatchesResponse{
				Round: RoundResponse{
					Stage:    match.Stage,
					Matchday: match.Matchday,
				},
				Matches: []MatchResponse{response},
			})
			lastRound = round
			continue
		}

		lastGroup := &groupedMatches[len(groupedMatches)-1]
		lastGroup.Matches = append(lastGroup.Matches, response)
	}

	return groupedMatches
}
