package api

import "strconv"

type GoalScorersResponse struct {
	Scorers []GoalScorer `json:"scorers"`
}

type GoalScorer struct {
	Player struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Nationality string `json:"nationality"`
	} `json:"player"`

	Team struct {
		ID int `json:"id"`
	} `json:"team"`

	Goals int `json:"goals"`

	Assists *int `json:"assists"`

	Penalties *int `json:"penalties"`
}

func (c *Client) FetchGoalScorers(competitionCode string, seasonYear int, limit int) ([]GoalScorer, error) {

	var response GoalScorersResponse

	//TODO chornic przed przecizeniem, za duzym limitem?
	if err := c.fetchWithQuery(
		"/competitions/"+competitionCode+"/scorers",
		&response,
		map[string]string{
			"season": strconv.Itoa(seasonYear),
			"limit":  strconv.Itoa(limit),
		},
	); err != nil {
		return nil, err
	}

	return response.Scorers, nil
}
