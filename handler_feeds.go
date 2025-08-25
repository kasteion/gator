package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error getting feeds: %w", err)
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("error getting feed user: %w", err)
		}

		fmt.Printf("*  Name: %s\n", feed.Name)
		fmt.Printf("   URL: %s\n", feed.Url)
		fmt.Printf("   User: %s\n", user.Name)
	}

	return nil
}