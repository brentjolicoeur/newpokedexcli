package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMapBack(cfg *config) error {
	if cfg.previous == nil || *cfg.previous == "" {
		fmt.Println("You're on the first page.")
		return nil
	}

	res, err := http.Get(*cfg.previous)
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
