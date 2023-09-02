package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/derekwilling/go-rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, struct{}{})
}

func handleErr(w http.ResponseWriter, r *http.Request) {
	responseWithError(w, http.StatusTeapot, "out of water")
}

func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	var params parameters
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Errorf("error parsing parameters: %w", err).Error())
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
	})
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, fmt.Errorf("error creating user: %w", err).Error())
	}

	respondWithJSON(w, http.StatusCreated, user)
}
