package main

import (
	"log"
	"os"

	"github.com/kasteion/gator/internal/config"
)

type state struct {
	config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	s := state{ config: &cfg }

	cmds := commands{
		registeredCommands: map[string]func(*state, command) error{},
	}

	cmds.register("login", handlerLogin)

	args := os.Args

	if len(args) < 2 {
		log.Fatal("missing arguments")
	}



	cmd := command{
		name: args[1],
		args: args[2:],
	}

	err = cmds.run(&s, cmd)
	if err != nil {
		log.Fatalf("%v", err)
	}
}