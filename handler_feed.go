package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/inileshsharma/rss-agg/internal/db"
	//"github.com/inileshsharma/rss-agg/internal/auth"
)

func (apicfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user db.User) {

	type parameters struct {
		Name string `json:"name"`
		Url string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondwithjson(w, 400, fmt.Sprintf("Error decoding request body: %v", err))
		return
	}

	feed, err := apicfg.DB.CreateFeed(r.Context(), db.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		Url: params.Url,
		UserID: user.ID,
	})
	if err != nil {
		respondwithjson(w, 400, fmt.Sprintf("couldn't create feed: %v", err))
		return
	}

	// fmt.Printf("TYPE: %T\n", UsertoUserResponse(user)) // should be main.User
    // fmt.Printf("VALUE: %+v\n", UsertoUserResponse(user))


	respondwithjson(w, 200, FeedtofeedResponse(feed))
}

func (apicfg *apiConfig) handlerGetFeed(w http.ResponseWriter, r *http.Request) {

	feed, err := apicfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondwithjson(w, 400, fmt.Sprintf("couldn't fetch feed: %v", err))
		return
	}

	respondwithjson(w, 200, FeedstofeedsResponse(feed))
}