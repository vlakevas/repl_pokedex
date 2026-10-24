package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocations(pageUrl *string) (ResLocations, error) {
	if pageUrl == nil {
		return ResLocations{}, fmt.Errorf("can't pass nil url pointer")
	}
	req, err := http.NewRequest("GET", *pageUrl, nil)
	if err != nil {
		return ResLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ResLocations{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return ResLocations{}, fmt.Errorf("bad status code: %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return ResLocations{}, err
	}

	var locationResponse ResLocations
	err = json.Unmarshal(data, &locationResponse)
	if err != nil {
		return ResLocations{}, err
	}
	return locationResponse, nil

}
