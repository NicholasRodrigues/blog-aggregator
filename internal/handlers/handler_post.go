package handlers

import (
	"github.com/NicholasRodrigues/blog-aggregator/internal/database"
	"github.com/NicholasRodrigues/blog-aggregator/pkg/jsonutil"
	"net/http"
	"strconv"
)

func (cfg *ApiConfig) HandlerGetPostFromUser(w http.ResponseWriter, r *http.Request, user database.User) {
	limitFromQuery := r.URL.Query().Get("limit")
	if limitFromQuery == "" {
		limitFromQuery = "10"
	}
	limitFromQueryInt, err := strconv.Atoi(limitFromQuery)
	if err != nil {
		jsonutil.RespondWithError(w, err, http.StatusBadRequest, "Invalid request payload")
		return
	}
	limit := 10
	if limitFromQuery != "" {
		limit = limitFromQueryInt
	}

	getPostsArg := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}

	posts, err := cfg.DB.GetPostsForUser(r.Context(), getPostsArg)

	if err != nil {
		jsonutil.RespondWithError(w, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	jsonutil.RespondWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}
