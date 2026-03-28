package main

import (
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Name of pokemon was not provided.")
		return nil
	}
	name := args[0]

	if pokemon, ok := cfg.caughtPokemon[name]; !ok {
		fmt.Println("You have not caught that pokemon.")
	} else {
		fmt.Printf("Name: %v\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, val := range pokemon.Types {
			fmt.Printf("  - %v\n", val.Type.Name)
		}
	}
	return nil
}
