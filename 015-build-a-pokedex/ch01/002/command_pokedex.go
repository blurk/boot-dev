package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("You haven't caught any pokemon yet!")
	}

	fmt.Println("Your Pokedex:")

	for pokemon := range cfg.caughtPokemon {
		fmt.Printf("- %s\n", pokemon)
	}

	return nil
}
