package http

import (
	"fmt"
	"github.com/NicholasRodrigues/blog-aggregator/internal/handlers"
	"github.com/NicholasRodrigues/blog-aggregator/internal/middlewares"
	"net/http"
	"os"
)

func StartServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	const filepathRoot = "."

	mux := http.NewServeMux()
	fsHandler := http.StripPrefix("/v1", http.FileServer(http.Dir(filepathRoot)))
	mux.Handle("/v1/", fsHandler)

	mux.HandleFunc("GET /v1/readiness", handlers.HandlerReadiness)
	mux.HandleFunc("GET /v1/err", handlers.HandlerErr)

	corsMux := middlewares.MiddlewareCors(mux)

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
