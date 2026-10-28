package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (IndividualPokemon, error) {
	pageUrl := baseUrl + "/pokemon/" + name

	data, exists := c.cache.Get(pageUrl)

	if !exists {
		req, err := http.NewRequest("GET", pageUrl, nil)
		if err != nil {
			return IndividualPokemon{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return IndividualPokemon{}, err
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			return IndividualPokemon{}, fmt.Errorf("bad status code: %d", res.StatusCode)
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return IndividualPokemon{}, err
		}

		c.cache.Add(pageUrl, data)
	}

	var poke IndividualPokemon
	err := json.Unmarshal(data, &poke)
	if err != nil {
		return IndividualPokemon{}, err
	}
	return poke, nil
}
