package routes

import (
	"database/sql"
	"github.com/NicholasRodrigues/blog-aggregator/internal/database"
	"github.com/NicholasRodrigues/blog-aggregator/internal/handlers"
	"net/http"
	"os"
)

func SetRoutes() (*http.ServeMux, *handlers.ApiConfig) {

	dbUrl := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, nil
	}

	dbQueries := database.New(db)

	apiCfg := handlers.ApiConfig{
		DB: dbQueries,
	}

	apiCfgPointer := &apiCfg

	const filepathRoot = "."
	mux := http.NewServeMux()
	fsHandler := http.StripPrefix("/v1", http.FileServer(http.Dir(filepathRoot)))
	mux.Handle("/v1/", fsHandler)

	mux.HandleFunc("GET /v1/readiness", handlers.HandlerReadiness)
	mux.HandleFunc("GET /v1/err", handlers.HandlerErr)

	mux.HandleFunc("POST /v1/users", apiCfg.HandlerCreateUser)
	mux.HandleFunc("GET /v1/users", apiCfg.HandlerGetUserByApiKey)

	mux.HandleFunc("POST /v1/feeds", apiCfg.AuthMiddleware(apiCfg.HandlerCreateFeed))
	mux.HandleFunc("POST /v1/feeds/follow", apiCfg.AuthMiddleware(apiCfg.HandlerFollowFeedCreate))
	mux.HandleFunc("DELETE /v1/feed_follows/{feedFollowID}", apiCfg.AuthMiddleware(apiCfg.HandlerFollowFeedDelete))
	mux.HandleFunc("GET /v1/feed_follows", apiCfg.AuthMiddleware(apiCfg.HandlerGetFeedFollowsByUserId))
	mux.HandleFunc("GET /v1/feeds", apiCfg.HandlerGetFeeds)

	mux.HandleFunc("GET /v1/posts", apiCfg.AuthMiddleware(apiCfg.HandlerGetPostFromUser))

	return mux, apiCfgPointer
}
