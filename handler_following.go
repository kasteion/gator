package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting user: %w", err)
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting feed_follows: %w", err)
	}

	for _, follow := range feedFollows {
		fmt.Printf("%v\n", follow)
	}

	return nil 
}
