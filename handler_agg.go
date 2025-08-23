package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	feed, err :=fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}

	fmt.Printf("Feed: %+v\n", feed)
	// printFeed(*feed)

	return  nil
}

// func printFeed(feed RSSFeed) {
// 	fmt.Println(feed.Channel.Title)
// 	fmt.Println(feed.Channel.Link)
// 	fmt.Println(feed.Channel.Description)
// 	for _, item := range feed.Channel.Item {
// 		fmt.Printf("  %s\n", item.Title)
// 		fmt.Printf("  %s\n", item.Link)
// 		fmt.Printf("  %s\n", item.Description)
// 		fmt.Printf("  %s\n", item.PubDate)
// 	}
// }