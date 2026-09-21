package database

import "context"

func (c *Client) ClearSeason(ctx context.Context, competitonID, startYear int) {

	c.DeleteSeasonMatches(ctx, competitonID, startYear)
	c.DeleteSeasonPlayers(ctx, competitonID, startYear)
}

func (c *Client) DeleteSeasonMatches(ctx context.Context, competitonID, startYear int) {
	filter := Filter{
		"competition_id":    competitonID,
		"start_season_year": startYear,
	}
	var goalScorers []SeasonPlayer
	c.db.WithContext(ctx).Where(filter).Delete(&goalScorers)
}

func (c *Client) DeleteSeasonPlayers(ctx context.Context, competitonID, startYear int) {
	filter := Filter{
		"competition_id":    competitonID,
		"start_season_year": startYear,
	}
	var goalScorers []SeasonPlayer
	c.db.WithContext(ctx).Where(filter).Delete(&goalScorers)
}
