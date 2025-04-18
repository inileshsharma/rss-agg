package main

import (
	"fmt"
	"net/http"

	"github.com/inileshsharma/rss-agg/internal/auth"
	"github.com/inileshsharma/rss-agg/internal/db"
)

type authedHandler func(http.ResponseWriter, *http.Request, db.User)

func (apicfg *apiConfig) authMiddleware(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetApiKeys(r.Header)
		if err != nil {
			respondwitherror(w,403,fmt.Sprintf("auth error: % v", err))
			return
		}
	
		user, err := apicfg.DB.GetUserByApikey(r.Context(), apikey)
		if err != nil {
			respondwitherror(w, 400, fmt.Sprintf("couldn't get user: %v", err))
			return
		}
	
		handler(w, r, user)
	}
}