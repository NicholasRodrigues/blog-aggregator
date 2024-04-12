package handlers

import (
	"github.com/NicholasRodrigues/blog-aggregator/internal/database"
	"github.com/NicholasRodrigues/blog-aggregator/pkg/jsonutil"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type Follow struct {
	ID        uuid.UUID `json:"id"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (cfg *ApiConfig) HandlerFollowFeedCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	params := parameters{}
	if err := jsonutil.DecodeRequestBody(r, &params); err != nil {
		jsonutil.RespondWithError(w, err, http.StatusBadRequest, "Invalid request payload")
		return
	}

	follow := Follow{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		FeedID:    params.FeedID,
		UserID:    user.ID,
		UpdatedAt: time.Now(),
	}

	followResponse, err := cfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams(follow))
	if err != nil {
		jsonutil.RespondWithError(w, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	jsonutil.RespondWithJSON(w, http.StatusCreated, followResponse)
}

func (cfg *ApiConfig) HandlerFollowFeedDelete(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := r.PathValue("feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		jsonutil.RespondWithError(w, err, http.StatusBadRequest, "Invalid feed follow id")
		return
	}

	err = cfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})

	if err != nil {
		jsonutil.RespondWithError(w, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	jsonutil.RespondWithJSON(w, http.StatusOK, nil)
}
