package handler

import (
	"fmt"
	"net/http"
)

const (
	CompetitionIDParam = "competitionID"
	StartYearParam     = "startYear"
	TeamIDParam        = "teamID"
	MatchIDParam       = "matchID"
	GoalScorerIDParam  = "goalScorerID"
	AreaIDParam        = "areaID"
)

func teamParam(r *http.Request) (int, error) {
	teamID, err := parseInt(r, TeamIDParam)
	if err != nil {
		return 0, fmt.Errorf("invalid team ID")
	}

	return teamID, nil
}

func competitionParam(r *http.Request) (int, error) {
	competitionID, err := parseInt(r, CompetitionIDParam)
	if err != nil {
		return 0, fmt.Errorf("invalid competition ID")
	}

	return competitionID, nil
}

func matchParam(r *http.Request) (int, error) {
	matchID, err := parseInt(r, MatchIDParam)
	if err != nil {
		return 0, fmt.Errorf("invalid match ID")
	}

	return matchID, nil
}

func goalScorerParam(r *http.Request) (int, error) {
	goalScorerID, err := parseInt(r, GoalScorerIDParam)
	if err != nil {
		return 0, fmt.Errorf("invalid goal scorer ID")
	}

	return goalScorerID, nil
}

func areaParam(r *http.Request) (int, error) {
	areaID, err := parseInt(r, AreaIDParam)
	if err != nil {
		return 0, fmt.Errorf("invalid area ID")
	}

	return areaID, nil
}

func editionParams(r *http.Request) (competitionID int, startYear int, err error) {
	competitionID, err = competitionParam(r)
	if err != nil {
		return 0, 0, err
	}

	startYear, err = parseInt(r, StartYearParam)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid start year")
	}

	return competitionID, startYear, nil
}
