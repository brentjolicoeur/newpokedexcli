package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var apiURL = "https://pokeapi.co/api/v2/location-area/"

type LocationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(cfg *config) error {
	var url string
	if cfg.next == nil {
		url = apiURL
	} else {
		url = *cfg.next
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var locations LocationAreaResponse
	if err := json.NewDecoder(res.Body).Decode(&locations); err != nil {
		return err
	}
	cfg.next, cfg.previous = &locations.Next, &locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
