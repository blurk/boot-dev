package main

import (
	"fmt"
	"time"
)

const url = "https://blog.boot.dev/index.xml"

func handleAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.name)
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenReqs)

	ticker := time.NewTicker(timeBetweenReqs)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		err := scrapeFeeds(s)

		if err != nil {
			return err
		}
	}
}
