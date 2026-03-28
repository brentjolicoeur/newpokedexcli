package main

import (
	"bufio"
	"fmt"
	"newpokedexcli/internal/pokeapi"
	"os"
	"strings"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			cleanedInput := cleanInput(input)
			if len(cleanedInput) == 0 {
				continue
			}
			commandName := cleanedInput[0]
			args := []string{}
			if len(cleanedInput) > 1 {
				args = cleanedInput[1:]
			}
			cmd, exists := getCommands()[commandName]
			if exists {
				err := cmd.callback(cfg, args)
				if err != nil {
					fmt.Println(err)
				}
				continue
			} else {
				fmt.Println("Unknown command")
				continue
			}
		} else {
			if scanner.Err() != nil {
				fmt.Print(scanner.Err())
			}
			break
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))

	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "List the pokemon found in a location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name",
			description: "Print info about a pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all captured pokemon",
			callback:    commandPokedex,
		},
	}
}
