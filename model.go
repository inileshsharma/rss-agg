package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/inileshsharma/rss-agg/internal/db"
)

type new_user struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
	ApiKey    string `json:"apikey"`
}

type feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
	Url    string `json:"apikey"`
	UserID    uuid.UUID `json:"user_id"`
}

func UsertoUserResponse(dbuser db.User) new_user {
	return new_user{
		ID:        dbuser.ID,
		CreatedAt: dbuser.CreatedAt,
		UpdatedAt: dbuser.UpdatedAt,
		Name:      dbuser.Name,
		ApiKey:    dbuser.ApiKey,
	}
}


func FeedtofeedResponse(dbuser db.Feed) feed {
	return feed{
		ID:        dbuser.ID,
		CreatedAt: dbuser.CreatedAt,
		UpdatedAt: dbuser.UpdatedAt,
		Name:      dbuser.Name,
		Url:	dbuser.Url,
		UserID:    dbuser.UserID,
	}
}

// from slice of feeds to slice of feeds response
func FeedstofeedsResponse(dbfeed []db.Feed) []feed {

	feeds := []feed{}
	for _, f := range dbfeed {
		feeds = append(feeds, FeedtofeedResponse(f))
	}
	return feeds

}

// feed follow response
type FeedFollows struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}
func FeedFollowResponse(dbfeedfollow db.FeedFollow) FeedFollows {
	return FeedFollows{
		ID:        dbfeedfollow.ID,
		CreatedAt: dbfeedfollow.CreatedAt,
		UpdatedAt: dbfeedfollow.UpdatedAt,
		UserID:    dbfeedfollow.UserID,
		FeedID:    dbfeedfollow.FeedID,
	}
}

func allfeedsFollowResponse(dbfeedfollow []db.FeedFollow) []FeedFollows {
	feeds := []FeedFollows{}
	for _, f := range dbfeedfollow {
		feeds = append(feeds, FeedFollowResponse(f))
	}
	return feeds
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func databasePostToPost(dbPost db.Post) Post {
	var description *string
	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}

	return Post{
		ID:          dbPost.ID,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
		Title:       dbPost.Title,
		Description: description,
		PublishedAt: dbPost.PublishedAt,
		Url:         dbPost.Url,
		FeedID:      dbPost.FeedID,
	}
}

func databasePostsToPosts(dbPosts []db.Post) []Post {
	posts := []Post{}
	for _, post := range dbPosts {
		posts = append(posts, databasePostToPost(post))
	}
	return posts
}