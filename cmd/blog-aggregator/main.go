package main

import (
	"github.com/NicholasRodrigues/blog-aggregator/internal/http"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	http.StartServer()
}
