package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))

	return words
}

func startREPL() {
	scanner := bufio.NewScanner(os.Stdin)

	var supportedCommands map[string]cliCommand
	supportedCommands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback: func() error {
				return commandHelp(supportedCommands)
			},
		},
	}

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			cleanedInput := cleanInput(input)
			if len(cleanedInput) == 0 {
				continue
			}
			firstWord := cleanedInput[0]
			if cmd, ok := supportedCommands[firstWord]; ok {
				err := cmd.callback()
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		} else {
			if scanner.Err() != nil {
				fmt.Print(scanner.Err())
			}
			break
		}
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp(commands map[string]cliCommand) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
