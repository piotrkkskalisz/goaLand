package api

import "strconv"

type TeamsResponse struct {
	Teams []Team `json:"teams"`
}

type Team struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	ShortName  string `json:"shortName"`
	TLA        string `json:"tla"`
	ClubColors string `json:"clubColors"`

	Area struct {
		ID int `json:"id"`
	} `json:"area"`

	Venue string `json:"venue"`
}

func (c *Client) FetchTeams(competitionCode string, seasonYear int) ([]Team, error) {

	var response TeamsResponse

	if err := c.fetchWithQuery(
		"/competitions/"+competitionCode+"/teams",
		&response,
		map[string]string{
			"season": strconv.Itoa(seasonYear),
		},
	); err != nil {
		return nil, err
	}

	return response.Teams, nil
}
