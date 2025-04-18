package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/inileshsharma/rss-agg/internal/db"
)

func startScaraping(db *db.Queries, concurrency int,
	timeBetweenReq time.Duration) {
	log.Printf("Scarping on %v foroutines every %s duration", concurrency, timeBetweenReq)
	ticker := time.NewTicker(timeBetweenReq)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedstoFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Printf("Error fetching feeds")
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go fetch(wg, feed, db)
		}
		wg.Wait()

	}
}

func fetch(wg *sync.WaitGroup, feed db.Feed, queries *db.Queries) {
	defer wg.Done()
	_, err := queries.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error while marking %s", feed.Name)
		return
	}
	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Printf("Error while fetching %s", feed.Name)
		return
	}
	for _, _feed := range rssFeed.Channel.Item {
		t, err := time.Parse(time.RFC1123Z, _feed.PubDate)
		if err != nil {
			log.Println("Error parsing date")
			continue
		}
		_, err = queries.CreatePost(context.Background(), db.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       _feed.Title,
			Description: sql.NullString{String: _feed.Description, Valid: true},
			PublishedAt: t,
			Url:         _feed.Link,
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Println("failed")
		}
	}
}


// func fetch(wg *sync.WaitGroup, feed db.Feed, db *db.Queries) {
// 	defer wg.Done()
// 	_, err := Queries.MarkFeedAsFetched(context.Background(), feed.ID)
// 	if err != nil {
// 		log.Printf("Error while marking %s", feed.Name)
// 		return
// 	}
// 	rssFeed, err := urlToFeed(feed.Url)
// 	if err != nil {
// 		log.Printf("Error while fetching %s", feed.Name)
// 		return
// 	}
// 	for _, _feed := range rssFeed.Channel.Item {
// 		t, err := time.Parse(time.RFC1123Z, _feed.PubDate)
// 		if err != nil {
// 			log.Println("Error parsing date")
// 			continue
// 		}
// 		_, err = db.CreatePost(context.Background(), db.CreatePostParams{
// 			ID:          uuid.New(),
// 			CreatedAt:   time.Now().UTC(),
// 			UpdatedAt:   time.Now().UTC(),
// 			Title:       _feed.Title,
// 			Description: sql.NullString{String: _feed.Description, Valid: true},
// 			PublishedAt: t,
// 			Url:         _feed.Link,
// 			FeedID:      feed.ID,
// 		})
// 		if err != nil {
// 			if strings.Contains(err.Error(), "duplicate key") {
// 				continue
// 			}
// 			log.Println("failed")
// 		}
// 	}
// }