package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ethanflaharty/PokedexCLI/internal/pokecache"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	cfg := &config{
		cache: pokecache.NewCache(5 * time.Second),
	}
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := words[1:]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	Next     *string
	Previous *string
	cache    *pokecache.Cache
	local    *location
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
			description: "Display 20 Pokemon locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the 20 previous Pokemon locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Gives a list of Pokemon in a given area",
			callback:    commandExplore,
		},
	}
}
