package main

import (
	"encoding/json"
	"learn/internal/auth"
	"learn/internal/database"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserID    uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) handleChiprsCreate(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)

	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error(), err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)

	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid JWT", err)
		return
	}

	type parameters struct {
		Body   string    `json:"body"`
		UserID uuid.UUID `json:"user_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	if len(userID) == 0 {
		respondWithError(w, http.StatusBadRequest, "user_id is empty", nil)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	profanes := []string{"kerfuffle", "sharbert", "fornax"}

	words := strings.Fields(params.Body)

	for i, word := range words {
		if slices.Contains(profanes, strings.ToLower(word)) {
			words[i] = "****"
		}
	}

	body := strings.Join(words, " ")

	chirp, err := cfg.db.CreateChirp(r.Context(), database.CreateChirpParams{
		Body:   body,
		UserID: userID,
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error(), err)
		return
	}

	respondWithJSON(w, http.StatusCreated, Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	})
}

func (cfg *apiConfig) handleChirpsGet(w http.ResponseWriter, r *http.Request) {
	data, err := cfg.db.GetChiprs(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error(), err)
	}

	chirps := make([]Chirp, len(data))

	for i, chirp := range chirps {
		chirps[i] = Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		}
	}

	respondWithJSON(w, http.StatusOK, chirps)
}

func (cfg *apiConfig) handleChirpsGetOne(w http.ResponseWriter, r *http.Request) {
	chirpID := r.PathValue("chirpID")

	parsedChirpID, _ := uuid.Parse(chirpID)

	data, err := cfg.db.GetChirp(r.Context(), parsedChirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Chirp not found", err)
		return
	}

	chirp := Chirp{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		Body:      data.Body,
		UserID:    data.UserID,
	}

	respondWithJSON(w, http.StatusOK, chirp)
}

func (cfg *apiConfig) handleChirpsDelete(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error(), err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid JWT", err)
		return
	}

	chirpID := r.PathValue("chirpID")
	parsedChirpID, err := uuid.Parse(chirpID)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error extracting ID", err)
		return
	}

	chirp, err := cfg.db.GetChirp(r.Context(), parsedChirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't find chirp", err)
		return
	}

	if chirp.UserID != userID {
		respondWithError(w, http.StatusForbidden, "You are not authorized to delete this chirp", err)
		return
	}

	err = cfg.db.DeleteChirp(r.Context(), parsedChirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't find chirp", err)
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}
