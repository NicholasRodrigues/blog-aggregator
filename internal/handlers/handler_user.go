package handlers

import (
	"github.com/NicholasRodrigues/blog-aggregator/internal/database"
	"github.com/NicholasRodrigues/blog-aggregator/pkg/jsonutil"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func (apiCfg *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	params := parameters{}
	if err := jsonutil.DecodeRequestBody(r, &params); err != nil {
		jsonutil.RespondWithError(w, err, http.StatusBadRequest, "Invalid request payload")
		return
	}

	userUuid := uuid.New()
	user := User{
		ID:        userUuid,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "Nicholas Rodrigues",
	}
	userResponse, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams(user))

	if err != nil {
		jsonutil.RespondWithError(w, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	jsonutil.RespondWithJSON(w, http.StatusCreated, userResponse)
}
