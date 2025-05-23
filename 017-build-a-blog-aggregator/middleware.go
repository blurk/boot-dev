package main

import (
	"context"

	"github.com/blurk/boot-dev/017-build-a-blog-aggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, c command) error {
		currentUser, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)

		if err != nil {
			return err
		}

		return handler(s, c, currentUser)
	}
}
