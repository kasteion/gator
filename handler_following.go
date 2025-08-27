package main

import (
	"context"
	"fmt"

	"github.com/kasteion/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting feed_follows: %w", err)
	}

	for _, follow := range feedFollows {
		fmt.Printf("%v\n", follow)
	}

	return nil 
}
