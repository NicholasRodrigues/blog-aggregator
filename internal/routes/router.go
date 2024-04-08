package routes

import (
	"database/sql"
	"github.com/NicholasRodrigues/blog-aggregator/internal/config"
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

	apiCfg := config.ApiConfig{
		DB: dbQueries,
	}

	const filepathRoot = "."
	mux := http.NewServeMux()
	fsHandler := http.StripPrefix("/v1", http.FileServer(http.Dir(filepathRoot)))
	mux.Handle("/v1/", fsHandler)

	mux.HandleFunc("GET /v1/readiness", handlers.HandlerReadiness)
	mux.HandleFunc("GET /v1/err", handlers.HandlerErr)

	return mux
}
