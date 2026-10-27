package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemons(location string) (ResPokemons, error) {
	pageUrl := baseUrl + "/location-area/" + location

	data, exists := c.cache.Get(pageUrl)
	if !exists {
		req, err := http.NewRequest("GET", pageUrl, nil)
		if err != nil {
			return ResPokemons{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return ResPokemons{}, err
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			return ResPokemons{}, fmt.Errorf("bad status code: %d", res.StatusCode)
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return ResPokemons{}, err
		}

		c.cache.Add(pageUrl, data)

	}
	var resPokemons ResPokemons
	err := json.Unmarshal(data, &resPokemons)
	if err != nil {
		return ResPokemons{}, err
	}
	return resPokemons, nil
}
