package database

import "context"

func (c *Client) GetEditionMatches(ctx context.Context, competitionID int, startYear int) ([]Match, error) {
	var matches []Match
	err := c.List(ctx, &matches, Filter{
		"competition_id":    competitionID,
		"start_season_year": startYear,
	},
	)
	return matches, err
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
