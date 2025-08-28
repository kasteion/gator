package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/kasteion/gator/internal/config"
	"github.com/kasteion/gator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	db *database.Queries
	config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	s := state{ config: &cfg }

	db, err := sql.Open("postgres", s.config.DBURL)
	if err != nil {
		log.Fatalf("error opening db connection: %v", err)
	}

	s.db = database.New(db)

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddfeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))

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