package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var locations RespShallowLocations

	if data, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(data, &locations); err != nil {
			return RespShallowLocations{}, err
		}
		return locations, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}
	if err = json.Unmarshal(body, &locations); err != nil {
		return RespShallowLocations{}, err
	}
	c.cache.Add(url, body)

	return locations, nil
}
func (c *Client) ListAreaInfo(locationName string) (LocationArea, error) {
	url := baseURL + "/location-area/" + locationName
	var locations LocationArea

	if data, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(data, &locations); err != nil {
			return LocationArea{}, err
		}
		return locations, nil
	}

	req, err := http.NewRequest("GET", url, nil)
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
	c.cache.Add(url, body)

	return locations, nil
}
