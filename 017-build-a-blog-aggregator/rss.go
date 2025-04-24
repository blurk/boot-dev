package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/blurk/boot-dev/017-build-a-blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func scrapeFeeds(s *state) error {
	feedToFetch, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	feed, err := fetchFeed(context.Background(), feedToFetch.Url)
	if err != nil {
		return err
	}

	err = s.db.MarkFeedFetched(context.Background(), feedToFetch.ID)
	if err != nil {
		return err
	}

	for _, item := range feed.Channel.Item {
		fmt.Printf("Saving posts of ... %v\n", feed.Channel.Title)
		err := savePost(s, item, feedToFetch.ID)
		fmt.Println("Done.")

		if err != nil {
			return err
		}
	}

	return nil
}

func savePost(s *state, post RSSItem, feedId uuid.UUID) error {
	pubDate, err := time.Parse(time.RFC1123Z, post.PubDate)

	if err != nil {
		return err
	}

	newPost, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
		ID:    uuid.New(),
		Title: post.Title,
		Description: sql.NullString{
			String: post.Description,
			Valid:  true,
		},
		Url: post.Link,
		PublishedAt: sql.NullTime{
			Time:  pubDate,
			Valid: true,
		},
		FeedID: feedId,
	})

	if err != nil {
		return err
	}

	fmt.Printf("%s was saved\n", newPost.Title)

	return nil
}
