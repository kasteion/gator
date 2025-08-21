package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <username>", cmd.name)
	}

	name := cmd.args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("error geting user: %v", err)
	}
	
	if err := s.config.SetUser(name); err != nil {
		return err
	}

	fmt.Println("user set successfully")
	return  nil
}