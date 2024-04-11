package routes

import (
	"database/sql"
	"github.com/NicholasRodrigues/blog-aggregator/internal/database"
	"github.com/NicholasRodrigues/blog-aggregator/internal/handlers"
	"net/http"
	"os"
)

func SetRoutes() *http.ServeMux {

	dbUrl := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil
	}

	dbQueries := database.New(db)

	apiCfg := handlers.ApiConfig{
		DB: dbQueries,
	}

	const filepathRoot = "."
	mux := http.NewServeMux()
	fsHandler := http.StripPrefix("/v1", http.FileServer(http.Dir(filepathRoot)))
	mux.Handle("/v1/", fsHandler)

	mux.HandleFunc("GET /v1/readiness", handlers.HandlerReadiness)
	mux.HandleFunc("GET /v1/err", handlers.HandlerErr)

	mux.HandleFunc("POST /v1/users", apiCfg.HandlerCreateUser)
	mux.HandleFunc("GET /v1/users", apiCfg.HandlerGetUserByApiKey)

	mux.Handle("POST /v1/feeds", apiCfg.AuthMiddleware(http.HandlerFunc(apiCfg.HandlerCreateFeed)))

	return mux
}
