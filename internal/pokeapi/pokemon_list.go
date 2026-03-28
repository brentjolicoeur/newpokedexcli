package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemonInfo(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name
	var pokemon Pokemon

	if data, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(data, &pokemon); err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	if err = json.Unmarshal(body, &pokemon); err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, body)

	return pokemon, nil
}
