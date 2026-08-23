package database

import (
	"backend/internal/utils"
	"cmp"
	"context"
	"slices"
	"time"
)

func (c *Client) GetEditionMatches(ctx context.Context, competitionID int, startYear int) ([]Match, error) {
	statuses := slices.Concat(utils.UpcomingMatchStatuses, utils.LiveMatchStatuses)

	return c.getEditionMatches(ctx, competitionID, startYear, statuses)
}

func (c *Client) GetEditionResult(ctx context.Context, competitionID int, startYear int) ([]Match, error) {
	matches, err := c.getEditionMatches(ctx, competitionID, startYear, utils.FinishedMatchStatuses)
	if err != nil {
		return nil, err
	}
	slices.Reverse(matches)

	return matches, nil
}

func (c *Client) getEditionMatches(ctx context.Context, competitionID int, startYear int,
	statuses []string) ([]Match, error) {
	var matches []Match
	if err := c.List(ctx, &matches, Filter{
		"competition_id":    competitionID,
		"start_season_year": startYear,
		"status":            statuses,
	}, preloadTeams...); err != nil {
		return nil, err
	}

	sortMatches(matches)
	return matches, nil
}

func (c *Client) FeaturedMatchesQuery(from time.Time, to time.Time) func(
	ctx context.Context, competitionID int, startYear int) ([]Match, error) {
	return func(ctx context.Context, competitionID int, startYear int) ([]Match, error) {
		var matches []Match
		err := c.buildQuery(ctx, preloadTeams).Where(Filter{
			"competition_id":    competitionID,
			"start_season_year": startYear,
			"status":            utils.DisplayableMatchStatuses,
		}).Where("start_time >= ? AND start_time < ?", from, to).Find(&matches).Error
		if err != nil {
			return nil, err
		}

		sortMatches(matches)
		return matches, nil
	}
}

func sortMatches(matches []Match) {
	slices.SortFunc(matches, func(a, b Match) int {
		if result := a.StartTime.Compare(b.StartTime); result != 0 {
			return result
		}

		return cmp.Compare(a.MatchID, b.MatchID)
	})
}

func (c *Client) GetTeamsMatches(ctx context.Context, teamID int) ([]Match, error) {
	var matches []Match

	err := c.ListOr(ctx, &matches, []Filter{
		{
			"home_team_id": teamID,
		}, {
			"away_team_id": teamID,
		},
	})

	if err != nil {
		return nil, err
	}

	return matches, err
}

func (c *Client) GetEditionGoalScorers(ctx context.Context, competitionID int, startYear int) ([]GoalScorer, error) {
	var goalScorers []GoalScorer
	err := c.List(ctx, &goalScorers, Filter{
		"competition_id":    competitionID,
		"start_season_year": startYear,
	},
	)
	return goalScorers, err
}
