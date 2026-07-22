package api

import "fmt"

type CompetitionsResponse struct {
	Competitions []Competition `json:"competitions"`
}

type Competition struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
	Type string `json:"type"`

	Area struct {
		ID int `json:"id"`
	} `json:"area"`
}

func (c *Client) FetchCompetitions() ([]Competition, error) {
	var response CompetitionsResponse

	if err := c.fetch("/competitions", &response); err != nil {
		return nil, err
	}

	return response.Competitions, nil
}

func (c *Client) FetchCompetition(code string) (Competition, error) {
	var response Competition

	url := fmt.Sprintf("/competitions/%s", code)
	if err := c.fetch(url, &response); err != nil {
		return Competition{}, err
	}

	return response, nil
}
