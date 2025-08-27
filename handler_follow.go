package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kasteion/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.name)
	}

	url := cmd.args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error getting feed: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(
		context.Background(), 
		database.CreateFeedFollowParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID: user.ID,
			FeedID: feed.ID,
		},
	)
	if err != nil {
		return fmt.Errorf("error creating feed_follow: %w", err)
	}
	fmt.Printf("%v\n", feedFollow)

	return nil
}