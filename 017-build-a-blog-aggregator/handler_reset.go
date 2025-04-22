package main

import (
	"context"
	"fmt"
)

func handleReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())

	if err != nil {
		return fmt.Errorf("couldn't delete all users: %w\n", err)
	}

	fmt.Println("Database reset successfully!")
	return nil
}
