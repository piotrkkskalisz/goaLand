package api

import "strconv"

type CompetitionDetailsResponse struct {
	Competition

	Season Edition `json:"season"`
}

type Edition struct {
	ID        int    `json:"id"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

func (c *Client) FetchEdition(competitionCode string, editionYear int) (Edition, error) {
	var response CompetitionDetailsResponse

	err := c.fetchWithQuery("/competitions/"+competitionCode+"/standings", &response,
		map[string]string{"season": strconv.Itoa(editionYear)})

	if err != nil {
		return Edition{}, err
	}

	return response.Season, nil
}
