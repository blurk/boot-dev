package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemon string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemon

	if value, exits := c.cache.Get(url); exits {

		var data Pokemon
		err := json.Unmarshal(value, &data)

		if err != nil {
			return Pokemon{}, err
		}

		return data, nil
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()
	if res.StatusCode > 299 {
		return Pokemon{}, errors.New("No pokemon found")
	}

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return Pokemon{}, nil
	}

	pokemonRes := Pokemon{}
	err = json.Unmarshal(data, &pokemonRes)

	return pokemonRes, nil
}
