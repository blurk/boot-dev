package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]

	if _, ok := cfg.caughtPokemon[name]; ok {
		return errors.New("Already caught this pokemon")
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...", name)

	rand.New(rand.NewSource(time.Now().UnixNano()))

	const maxBaseExp = 300.0

	// Calculate catch chance: higher baseExp means lower chance
	catchChance := max((maxBaseExp-float64(pokemon.BaseExperience))/maxBaseExp, 0)
	roll := rand.Float64()

	if roll < catchChance {
		fmt.Printf("%s escaped!\n", name)
	} else {
		cfg.caughtPokemon[name] = pokemon
		fmt.Printf("%s was caught!\n", name)
	}

	return nil
}
