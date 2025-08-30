package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/kasteion/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage %s <time_between_reqs>", cmd.name)
	}

	duration, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("error parsing duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %v\n", duration)

	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		// fmt.Println("#########################")
		// feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
		// if err != nil {
			// return fmt.Errorf("error fetching feed: %w", err)
		// }

		// fmt.Printf("Feed: %+v\n", feed)
		scrapeFeeds(context.Background(), s)
	}

	// printFeed(*feed)
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

func scrapeFeeds(ctx context.Context, s *state) {
	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return 
	}

	err = s.db.MarkFeedFetched(ctx, 
		database.MarkFeedFetchedParams{
			ID: feed.ID,
			UpdatedAt: time.Now(),
		},
	)
	if err != nil {
		return 
	}

	rssFeed, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		return 
	}

	// fmt.Println(rssFeed.Channel.Title)
	for _, item := range rssFeed.Channel.Item {
		publishedAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Println(err)
		}

		_, err = s.db.CreatePost(
			ctx,
			database.CreatePostParams{
				ID: uuid.New(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Title: item.Title,
				Url: item.Link,
				Description: item.Description,
				PublishedAt: publishedAt,
				FeedID: feed.ID,
			},
		)

		if err != nil && strings.Compare("pq: duplicate key value violates unique constraint \"posts_url_key\"", err.Error()) != 0 {
			log.Println(err)
			
		}
	}
}