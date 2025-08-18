package main

import (
	"fmt"
	"log"

	"github.com/kasteion/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	if err := cfg.SetUser("kasteion"); err != nil {
		log.Fatalf("couldn't set current user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	// fmt.Println(*cfg.DBUrl)
	// fmt.Println(*cfg.CurrentUserName)
	fmt.Printf("Read config: %+v\n", cfg)
}