package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) ListPokemons(location string) (Location, error) {
	url := baseURL + "/location-area/" + location

	if val, ok := c.cache.Get(url); ok {
		pokemonsResp := Location{}
		err := json.Unmarshal(val, &pokemonsResp)

		if err != nil {
			return Location{}, err
		}

		return pokemonsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, nil
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return Location{}, errors.New("No pokemon found")
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	locationRes := Location{}
	err = json.Unmarshal(data, &locationRes)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, data)

	return locationRes, nil
}
