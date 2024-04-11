package handlers

import (
	"context"
	"github.com/NicholasRodrigues/blog-aggregator/internal/auth"
	"net/http"
)

//type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *ApiConfig) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "user", user)))
	})
}
