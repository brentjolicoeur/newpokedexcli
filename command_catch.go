package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("Name of pokemon to be caught was not provided.")
	}
	name := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := cfg.pokeapiClient.ListPokemonInfo(name)
	if err != nil {
		return err
	}

	difficulty := float64(pokemon.BaseExperience) / 600.0
	catchChance := float64(1.0 - difficulty)

	if catchChance < 0.02 {
		catchChance = 0.02
	}
	if catchChance > rand.ExpFloat64() {
		fmt.Printf("%s was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
