package main

import (
	"log"
	"os"

	"github.com/blurk/boot-dev/017-build-a-blog-aggregator/internal/config"
)

type state struct {
	config *config.Config
}

func main() {
	config, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := &state{
		config: &config,
	}

	commands := commands{
		commands: make(map[string]func(*state, command) error),
	}
	commands.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = commands.run(programState, command{name: cmdName, args: cmdArgs})

	if err != nil {
		log.Fatal(err)
	}
}
