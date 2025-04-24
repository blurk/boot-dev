package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/blurk/boot-dev/017-build-a-blog-aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, currentUser database.User) error {
	var limit int64 = 2
	var err error

	if len(cmd.args) == 1 {
		limit, err = strconv.ParseInt(cmd.args[0], 10, 32)

		if err != nil {
			return err
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: currentUser.ID,
		Limit:  int32(limit),
	})

	if err != nil {
		return err
	}

	fmt.Println("Your post:")
	for _, post := range posts {
		fmt.Printf("* Title:          %v\n", post.Title)
		fmt.Printf("* Description:    %v\n", post.Description.String)
		fmt.Printf("* Link:           %v\n", post.Url)
		fmt.Println()
	}

	return nil
}
