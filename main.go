package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)


func main(){

	// fmt.Print("Hello, World!\n")

	godotenv.Load(".env")

	portstring := os.Getenv("PORT")

	if portstring == "" {
		log.Fatal("PORT environment variable not set")
	}

	fmt.Println("PORT environment variable is set to:", portstring)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*","http://*"},	// Allow from all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},	
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		MaxAge:           300, // Maximum value not ignored by any of major browsers
		AllowCredentials: false, // Allows cookies to be sent
		}))
	
		
	v1router := chi.NewRouter()
	v1router.Get("/healthz", handlerReadiness)

	v1router.Get("/error", errorHandler)
	
	router.Mount("/v1", v1router)

	srv := &http.Server{
		Addr:    ":" + portstring,
		Handler: router,
	}

	log.Printf("Starting server on port %v", portstring)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
