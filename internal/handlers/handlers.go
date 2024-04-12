package handlers

import (
	"github.com/NicholasRodrigues/blog-aggregator/internal/auth"
	"github.com/NicholasRodrigues/blog-aggregator/internal/database"
	"net/http"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *ApiConfig) AuthMiddleware(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetApiKey(r.Header)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if apikey == "" {
			http.Error(w, "API key is required", http.StatusUnauthorized)
			return
		}
		user, err := cfg.DB.GetUserByApiKey(r.Context(), apikey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		handler(w, r, user)
	}
}
