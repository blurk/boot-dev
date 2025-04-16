package main

import (
	"time"

	"github.com/blurk/boot-dev/015-build-a-pokedex/ch01/002/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)

	config := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(config)
}
