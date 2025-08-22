package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if err := s.db.ResetUsers(context.Background()); err != nil {
		return fmt.Errorf("error reseting users: %w", err)
	}

	fmt.Println("users reset successfully")

	return nil
}