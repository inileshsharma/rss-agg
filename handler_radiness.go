package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondwithjson(w,200,struct{}{})

}