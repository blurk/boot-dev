package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/blurk/boot-dev/017-build-a-blog-aggregator/internal/config"
	"github.com/blurk/boot-dev/017-build-a-blog-aggregator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	db     *database.Queries
	config *config.Config
}

func main() {
	config, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", config.DBUrl)

	if err != nil {
		log.Fatalf("error connect to database: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	programState := &state{
		config: &config,
		db:     dbQueries,
	}

	commands := commands{
		commands: make(map[string]func(*state, command) error),
	}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handleReset)
	commands.register("users", handlerUsers)

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
