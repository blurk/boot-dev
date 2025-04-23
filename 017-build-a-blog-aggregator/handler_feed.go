package main

import (
	"context"
	"fmt"

	"github.com/blurk/boot-dev/017-build-a-blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, currentUser database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %s <name of feed> <url of feed>", cmd.name)
	}

	feedName := cmd.args[0]
	feedUrl := cmd.args[1]

	newFeed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:     uuid.New(),
		Name:   feedName,
		Url:    feedUrl,
		UserID: currentUser.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't add new feed: %w", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{ID: uuid.New(),
		UserID: currentUser.ID,
		FeedID: newFeed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println("Feed created successfully:")
	printFeed(newFeed)
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())

	if err != nil {
		return fmt.Errorf("couldn't add new feed: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	fmt.Printf("Found %d feeds:\n", len(feeds))
	printFeeds(feeds)
	fmt.Println("=====================================")

	return nil
}

func printFeeds(feeds []database.GetFeedsRow) {
	for _, feed := range feeds {
		fmt.Printf("* Name:       %s\n", feed.Name)
		fmt.Printf("* Url:        %s\n", feed.Url)
		fmt.Printf("* User:       %s\n", feed.Username)
		fmt.Println()
	}
}

func handlerFollow(s *state, cmd command, currentUser database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <url of feed>", cmd.name)
	}

	url := cmd.args[0]

	feed, err := s.db.GetFeed(context.Background(), url)

	if err != nil {
		return err
	}

	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:     uuid.New(),
		UserID: currentUser.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		return fmt.Errorf("couldn't add new feed: %w", err)
	}

	fmt.Println("Feed followed:")
	fmt.Printf("* Feed:        %v\n", follow.FeedName)
	fmt.Printf("* User:        %v\n", follow.UserName)
	fmt.Println()

	return nil
}

func handlerFollowing(s *state, cmd command, currentUser database.User) error {
	following, err := s.db.GetFeedFollowForUser(context.Background(), currentUser.ID)

	if err != nil {
		return err
	}

	if len(following) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	fmt.Printf("%s following feed:\n", currentUser.Name)
	for _, feed := range following {
		fmt.Printf("* Name:         %+v\n", feed.Feedname)
		fmt.Printf("* Creator:      %+v\n", feed.Creator)
		fmt.Println()
	}

	return nil
}

func handlerUnfollow(s *state, cmd command, currentUser database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <url of feed>", cmd.name)
	}

	url := cmd.args[0]

	err := s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: currentUser.ID,
		Url:    url,
	})

	if err != nil {
		return err
	}

	fmt.Println("Unfollowing successfully")

	return nil
}
