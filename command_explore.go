package main

import (
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Location to explore was not provided.")
		return nil
	}
	location := args[0]
	fmt.Printf("Exploring %s\n", location)
	url := "https://pokeapi.co/api/v2/location-area/" + location

	areaInfo, err := cfg.pokeapiClient.ListAreaInfo(url)
	if err != nil {
		return err
	}
	for _, encounter := range areaInfo.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
