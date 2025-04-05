package main

import (
	"encoding/json"
	"net/http"
)

func respondwithjson(w http.ResponseWriter, code int, payload interface{}) {

	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondwitherror(w http.ResponseWriter, code int, message string) {
	
	if code >= 500 {
		message = "5XX Server Error"
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	respondwithjson(w, code, errorResponse{Error: message})
}
