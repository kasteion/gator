package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/kasteion/gator/internal/database"
)

func handlerAddfeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.name)
	}

	name := cmd.args[0]
	url := cmd.args[1]

	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting user: %w", err)
	}

	feed, err := s.db.CreateFeed(
		context.Background(), 
		database.CreateFeedParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name: name,
			Url: url,
			UserID: user.ID,
		},
	)
	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}

	_, err = s.db.CreateFeedFollow(
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
		return fmt.Errorf("error created feed_follow: %w", err)
	}

	fmt.Println("feed created successfully")
	log.Printf("feed: %v\n", feed)

	return  nil
}