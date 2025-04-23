package main

import (
	"context"
	"fmt"
)

func scrapeFeeds(s *state) error {
	nextFeedToFetch, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	feed, err := fetchFeed(context.Background(), nextFeedToFetch.Url)
	if err != nil {
		return err
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeedToFetch.ID)
	if err != nil {
		return err
	}

	fmt.Printf("Item of: %s\n", feed.Channel.Title)
	for i, item := range feed.Channel.Item {
		fmt.Printf("%d: %s\n", i, item.Title)
	}

	return nil
}
