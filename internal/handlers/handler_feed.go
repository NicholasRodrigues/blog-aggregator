package handlers

import (
	"github.com/NicholasRodrigues/blog-aggregator/internal/database"
	"github.com/NicholasRodrigues/blog-aggregator/pkg/jsonutil"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (cfg *ApiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	params := parameters{}
	if err := jsonutil.DecodeRequestBody(r, &params); err != nil {
		jsonutil.RespondWithError(w, err, http.StatusBadRequest, "Invalid request payload")
		return
	}

	feedUuid := uuid.New()
	feed := Feed{
		ID:        feedUuid,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	}

	feedResponse, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams(feed))
	if err != nil {
		jsonutil.RespondWithError(w, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	feedFollow := Follow{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		FeedID:    feedUuid,
		UserID:    user.ID,
		UpdatedAt: time.Now(),
	}
	feedFollowResponse, err := cfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams(feedFollow))
	if err != nil {
		jsonutil.RespondWithError(w, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	jsonutil.RespondWithJSON(w, http.StatusCreated, struct {
		feed       Feed
		feedFollow FeedFollow
	}{
		feed:       databaseFeedToFeed(feedResponse),
		feedFollow: databaseFeedFollowToFeedFollow(feedFollowResponse),
	})
}

func (cfg *ApiConfig) HandlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetFeeds(r.Context())

	if err != nil {
		jsonutil.RespondWithError(w, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	jsonutil.RespondWithJSON(w, http.StatusOK, feeds)
}
