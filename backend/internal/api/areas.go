package api

type AreasResponse struct {
	Areas []Area `json:"areas"`
}

type Area struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	CountryCode string  `json:"countryCode"`
	ParentArea  *string `json:"parentArea"`
}

func (c *Client) FetchAreas() ([]Area, error) {
	var response AreasResponse

	if err := c.fetch("/areas", &response); err != nil {
		return nil, err
	}

	return response.Areas, nil
}
