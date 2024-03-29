package handlers

import (
	"github.com/NicholasRodrigues/blog-aggregator/pkg/jsonutil"
	"net/http"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	jsonutil.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func HandlerErr(w http.ResponseWriter, r *http.Request) {
	jsonutil.RespondWithError(w, nil, http.StatusInternalServerError, "Internal Server Error")
}
