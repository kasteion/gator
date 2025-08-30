package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/kasteion/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	
	if len(cmd.args) == 1 {
		l, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("usage: %s <limit> - limit should be a number", cmd.name)
		}

		limit = l
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{ UserID: user.ID, Limit: int32(limit)})
	if err != nil {
		return fmt.Errorf("error getting posts: %w", err)
	}

	for _, post := range posts {
		fmt.Printf("%v\n", post)
	}

	return nil
}
