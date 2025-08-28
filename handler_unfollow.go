package main

import (
	"context"
	"fmt"

	"github.com/kasteion/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.name)
	}

	url := cmd.args[0]

	if err := s.db.DeleteByUserURL(context.Background(), database.DeleteByUserURLParams{ UserID: user.ID, Url: url }); err != nil {
		return fmt.Errorf("error deleting feed follow: %w", err)
	}

	fmt.Println("feed unfollowed successfully")

	return  nil
}