package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListAreaInfo(pageURL string) (LocationArea, error) {

	var locations LocationArea

	if data, ok := c.cache.Get(pageURL); ok {
		if err := json.Unmarshal(data, &locations); err != nil {
			return LocationArea{}, err
		}
		return locations, nil
	}

	req, err := http.NewRequest("GET", pageURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}
	if err = json.Unmarshal(body, &locations); err != nil {
		return LocationArea{}, err
	}
	c.cache.Add(pageURL, body)

	return locations, nil
}
