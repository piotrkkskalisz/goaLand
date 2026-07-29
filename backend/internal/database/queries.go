package database

import "context"

func (c *Client) GetEditionMatches(ctx context.Context, competitionID int, startYear int) ([]Match, error) {
	var matches []Match
	err := c.Get(ctx, &matches, Filter{
		"competition_id": competitionID,
		"start_year":     startYear,
	},
	)
	return matches, err
}

func (c *Client) GetTeamsMatches(ctx context.Context, teamID int) ([]Match, error) {
	var matches []Match
	err := c.Get(ctx, &matches, Filter{
		"team_id": teamID,
	},
	)
	return matches, err
}

func (c *Client) GetEditionGoalScorers(ctx context.Context, competitionID int, startYear int) ([]GoalScorer, error) {
	var goalScorers []GoalScorer
	err := c.Get(ctx, &goalScorers, Filter{
		"competition_id":    competitionID,
		"start_season_year": startYear,
	},
	)
	return goalScorers, err
}
