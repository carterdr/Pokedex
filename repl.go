package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/carterdr/Pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}
type config struct {
	pokeApiClient pokeapi.Client
	nextUrl       *string
	previousUrl   *string
	caught        map[string]pokeapi.Pokemon
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
			description: "Displays Next 20 Locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays Previous 20 Locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location name>",
			description: "Explores a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon name>",
			description: "Attempts to catch pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "View details about a caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "See all the pokemon you've caught",
			callback:    commandPokedex,
		},
	}
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	cliMap := getCommands()
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := cliMap[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Invalid command. Type 'help' for available commands.")
			continue
		}
	}
}

func cleanInput(text string) (output []string) {
	text = strings.ToLower(text)
	output = strings.Fields(text)
	return output
}
