package main

import (
	"fmt"
)

func commandMap(cfg *config, args []string) error {
	locations, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
