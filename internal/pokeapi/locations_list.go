package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) LocationList(pageURL *string) (LocationAreas, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return LocationAreas{}, err
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return LocationAreas{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	locationResp := LocationAreas{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return LocationAreas{}, err
	}

	return locationResp, nil
}