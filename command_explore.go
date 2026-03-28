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

	areaInfo, err := cfg.pokeapiClient.ListAreaInfo(location)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, encounter := range areaInfo.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
