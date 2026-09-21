package sync

import (
	"context"
	"errors"
)

var live_leagues = []string{"PL", "BL1"}

func (s *Sync) updateEdition(ctx context.Context, target SeasonTarget) error {
	competitionID, ok := s.seasons.CompetitionID(target.CompetitionCode)

	if !ok {
		return errors.New("not found competition ID")
	}
	season := Season{
		CompetitionID:   competitionID,
		CompetitionCode: target.CompetitionCode,
		StartYear:       target.StartYear,
	}

	s.databaseClient.DeleteSeasonPlayers(ctx, season.CompetitionID, season.StartYear)

	apiTeams, err := s.apiClient.FetchTeams(season.CompetitionCode, season.StartYear)
	if err != nil {
		return err
	}

	err = s.initPlayers(ctx, apiTeams, season)
	if err != nil {
		return err
	}

	change, err := s.initMatches(ctx, season)
	if err != nil {
		return err
	}
	if change {
		s.storage.RefreshTable(ctx, season.CompetitionID, season.StartYear)
	}

	if err := s.initGoalScorers(ctx, season, defaultGoalScorerLimit); err != nil {
		return err
	}
	return nil
}
