package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error getting users: %w", err)
	}

	for _, user := range users {
		fmtd := fmt.Sprintf("* %s", user.Name)
		if s.config.CurrentUserName == user.Name {
			fmtd += " (current)"
		}
		fmt.Println(fmtd)
	}

	return nil
}