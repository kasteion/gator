package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/kasteion/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <username>", cmd.name)
	}

	name := cmd.args[0]

	user, err := s.db.CreateUser(
		context.Background(), 
		database.CreateUserParams{ 
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name: name,
		},
	)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	if err := s.config.SetUser(name); err != nil {
		return err
	}

	fmt.Println("user created successfully")
	log.Printf("user: %v\n", user)

	return nil
}