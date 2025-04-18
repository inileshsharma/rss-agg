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

func (apicfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondwithjson(w, 400, fmt.Sprintf("Error decoding request body: %v", err))
		return
	}

	user, err := apicfg.DB.CreateUser(r.Context(), db.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err != nil {
		respondwithjson(w, 400, fmt.Sprintf("couldn't create user: %v", err))
		return
	}

	// fmt.Printf("TYPE: %T\n", UsertoUserResponse(user)) // should be main.User
    // fmt.Printf("VALUE: %+v\n", UsertoUserResponse(user))


	respondwithjson(w, 200, UsertoUserResponse(user))
}

func (apicfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user db.User) {

	respondwithjson(w, 200, UsertoUserResponse(user))

}

func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user db.User) {
	posts, err := apiCfg.DB.GetPostForUser(r.Context(), db.GetPostForUserParams{
		UserID: user.ID,
		Limit:  10,
	})

	if err != nil {
		respondwitherror(w, 400, fmt.Sprintf("Couldn't get posts: %v", err))
		return
	}

	respondwithjson(w, 200, databasePostsToPosts(posts))
}