package main

import "net/http"

func errorHandler(w http.ResponseWriter, r *http.Request) {
	respondwitherror(w, 400, "Internal Server Error")
}