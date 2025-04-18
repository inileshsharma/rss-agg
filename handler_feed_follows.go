package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/inileshsharma/rss-agg/internal/db"
	//"github.com/inileshsharma/rss-agg/internal/auth"
)

func (apicfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user db.User) {

	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondwithjson(w, 400, fmt.Sprintf("Error decoding request body: %v", err))
		return
	}

	feedfollow, err := apicfg.DB.CreateFeedFollows(r.Context(), db.CreateFeedFollowsParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: params.FeedId,

	})
	if err != nil {
		respondwithjson(w, 400, fmt.Sprintf("couldn't create feed follows: %v", err))
		return
	}

	// fmt.Printf("TYPE: %T\n", UsertoUserResponse(user)) // should be main.User
    // fmt.Printf("VALUE: %+v\n", UsertoUserResponse(user))


	respondwithjson(w, 200, FeedFollowResponse(feedfollow))
}

func (apicfg *apiConfig) handlerGetFeedFollow(w http.ResponseWriter, r *http.Request, user db.User) {

	allfeeds, err := apicfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondwithjson(w, 400, fmt.Sprintf("couldn't get feed following : %v", err))
		return
	}

	// fmt.Printf("TYPE: %T\n", UsertoUserResponse(user)) // should be main.User
    // fmt.Printf("VALUE: %+v\n", UsertoUserResponse(user))


	respondwithjson(w, 200, allfeedsFollowResponse(allfeeds))
}

func (apicfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user db.User) {

	feedfollowidstr := chi.URLParam(r, "feed_id")
	feedfollowid, err := uuid.Parse(feedfollowidstr)

	if err != nil {
		respondwithjson(w, 400, fmt.Sprintf("couldn't parse feed follow id: %v", err))
		return
	}

	err = apicfg.DB.DeleteFeedFollows(r.Context(), db.DeleteFeedFollowsParams{
		ID: feedfollowid,
		UserID: user.ID,
	})
	if err != nil {
		respondwithjson(w, 400, fmt.Sprintf("couldn't unfollow feed id: %v", err))
		return
	}

	respondwithjson(w, 200, struct{}{})

}
