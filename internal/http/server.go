package http

import (
	"fmt"
	"github.com/NicholasRodrigues/blog-aggregator/internal/middlewares"
	"github.com/NicholasRodrigues/blog-aggregator/internal/routes"
	"net/http"
	"os"
)

func StartServer() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	mux := routes.SetRoutes()
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
