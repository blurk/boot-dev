package main

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/blurk/boot-dev/015-build-a-pokedex/ch01/002/internal/pokeapi"
	"github.com/chzyer/readline"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
}

func startRepl(config *config) {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          "Pokedex > ",
		HistoryFile:     "/tmp/pokedex_history.tmp",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		log.Fatalf("Failed to initialize readline: %v", err)
	}

	// Ensure the readline instance is closed when main exits.
	// This is important for restoring terminal settings.
	defer rl.Close()

	fmt.Println("Welcome to Pokedex. Type 'exit' or press Ctrl+D to quit.")

	// Main loop to read commands
	for {
		// Read a line of input from the user.
		// Readline handles history navigation (up/down arrows) and line editing.
		line, err := rl.Readline()

		// Handle potential errors during input reading.
		if err == readline.ErrInterrupt {
			// User pressed Ctrl+C
			// If line was empty, exit. Otherwise, clear current line.
			if len(line) == 0 {
				break // Exit the loop
			} else {
				continue // Continue to next iteration, clearing input
			}
		} else if err == io.EOF {
			// User pressed Ctrl+D
			break // Exit the loop
		} else if err != nil {
			// Other potential errors
			log.Printf("Error reading line: %v", err)
			break // Exit on other errors
		}

		// Trim whitespace from the input.
		words := strings.Fields(strings.ToLower(line))

		commandName := words[0]

		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		// --- Process the command ---
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(config, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

		// The command is automatically added to history by readline *after*
		// Readline() returns successfully.
	}

	commandExit(config)
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inpsect <pokemon_name>",
			description: "Inpsect your caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Show your pokedex",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
