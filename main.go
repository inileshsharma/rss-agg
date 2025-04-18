package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/inileshsharma/rss-agg/internal/db"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *db.Queries
	// cfg field removed as it is not defined or used
}

func main() {

	// fmt.Print("Hello, World!\n")

	godotenv.Load(".env")

	portstring := os.Getenv("PORT")

	if portstring == "" {
		log.Fatal("PORT environment variable not set")
	}
	fmt.Println("PORT environment variable is set to:", portstring)

	db_url := os.Getenv("DB_URL")
	if db_url == "" {
		log.Fatal("database URL not set")
	}

	conn, err := sql.Open("postgres", db_url)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
    db := db.New(conn)
	apicgf := apiConfig{
		DB: db,
	}

	go startScaraping(
		db, 10, time.Minute,
	)


	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"}, // Allow from all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		MaxAge:           300,   // Maximum value not ignored by any of major browsers
		AllowCredentials: false, // Allows cookies to be sent
	}))

	v1router := chi.NewRouter()
	v1router.Get("/healthz", handlerReadiness)
	v1router.Get("/error", errorHandler)
	v1router.Post("/users", apicgf.handlerCreateUser)
	v1router.Get("/users", apicgf.authMiddleware(apicgf.handlerGetUser))
    v1router.Post("/feeds", apicgf.authMiddleware(apicgf.handlerCreateFeed))
	v1router.Get("/feeds", apicgf.handlerGetFeed)
	v1router.Post("/feedfollows", apicgf.authMiddleware(apicgf.handlerCreateFeedFollow))
	v1router.Get("/feedfollows", apicgf.authMiddleware(apicgf.handlerGetFeedFollow))
	v1router.Delete("/feedfollows/{feed_id}", apicgf.authMiddleware(apicgf.handlerDeleteFeedFollow))

	v1router.Get("/posts", apicgf.authMiddleware(apicgf.handlerGetPostsForUser))
	router.Mount("/v1", v1router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portstring,
	}

	log.Printf("Server starting on port %v", portstring)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
