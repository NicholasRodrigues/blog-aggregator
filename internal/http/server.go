package http

import (
	"fmt"
	"github.com/NicholasRodrigues/blog-aggregator/internal/middlewares"
	"github.com/NicholasRodrigues/blog-aggregator/internal/pkg"
	"github.com/NicholasRodrigues/blog-aggregator/internal/routes"
	"log"
	"net/http"
	"os"
	"time"
)

func StartServer() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	mux, apiConfigPointer := routes.SetRoutes()
	corsMux := middlewares.MiddlewareCors(mux)

	dbQueries := apiConfigPointer.DB
	const collectionConcurrency = 10
	const collectionInterval = time.Minute
	log.Println("Starting scraping...")
	go pkg.StartScraping(dbQueries, collectionConcurrency, collectionInterval)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}
	fmt.Println("Server created at: " + "http://localhost:" + port + "/app/")

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}

}
