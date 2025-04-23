package main

import (
	"context"
	"fmt"
)

const url = "https://www.wagslane.dev/index.xml"

func handleAgg(s *state, cmd command) error {
	res, err := fetchFeed(context.Background(), url)

	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w\n", err)
	}

	fmt.Printf("Feed: %+v\n", res)

	return nil
}
