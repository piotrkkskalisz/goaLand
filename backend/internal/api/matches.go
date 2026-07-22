package api

import "strconv"

type MatchesResponse struct {
	Matches []Match `json:"matches"`
}

type Match struct {
	ID int `json:"id"`

	UtcDate  string `json:"utcDate"`
	Status   string `json:"status"`
	Matchday int    `json:"matchday"`

	HomeTeam struct {
		ID int `json:"id"`
	} `json:"homeTeam"`

	AwayTeam struct {
		ID int `json:"id"`
	} `json:"awayTeam"`

	Score struct {
		FullTime struct {
			Home *int `json:"home"`
			Away *int `json:"away"`
		} `json:"fullTime"`

		HalfTime struct {
			Home *int `json:"home"`
			Away *int `json:"away"`
		} `json:"halfTime"`
	} `json:"score"`
}

func (c *Client) FetchMatches(competitionCode string, seasonYear int) ([]Match, error) {

	var response MatchesResponse

	if err := c.fetchWithQuery(
		"/competitions/"+competitionCode+"/matches",
		&response,
		map[string]string{
			"season": strconv.Itoa(seasonYear),
		},
	); err != nil {
		return nil, err
	}

	return response.Matches, nil
}
