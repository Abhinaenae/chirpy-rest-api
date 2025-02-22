package main

import (
	"chirpy/internal/auth"
	"chirpy/internal/database"
	"chirpy/internal/filter"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type chirpReq struct {
	Body string `json:"body"`
	//UserID uuid.UUID `json:"user_id"`
	//Felt redudant
}

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserID    uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) postChirpHandler(w http.ResponseWriter, r *http.Request) {

	tokenString, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	userID, err := auth.ValidateJWT(tokenString, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	decoder := json.NewDecoder(r.Body)
	chirp := chirpReq{}
	if err := decoder.Decode(&chirp); err != nil {
		log.Printf("Error decoding JSON: %s", err)
		respondWithError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	// Ensure the user making the request is the same as the user in the chirp
	/*
		if userID != chirp.UserID {
			respondWithError(w, http.StatusUnauthorized, "Not matched up with correct user")
			return
		}
	*/

	if len(chirp.Body) > 140 {
		respondWithError(w, http.StatusBadRequest, "chirp length is too long")
		return
	}

	cleaned := filter.FilterProfanity(chirp.Body)
	chirp.Body = cleaned

	savedChirp, err := cfg.dbQueries.CreateChirp(r.Context(), database.CreateChirpParams{
		Body:   chirp.Body,
		UserID: userID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to post chirp")
		return
	}

	chirpRes := Chirp{
		ID:        savedChirp.ID,
		CreatedAt: savedChirp.CreatedAt,
		UpdatedAt: savedChirp.UpdatedAt,
		Body:      savedChirp.Body,
		UserID:    savedChirp.UserID,
	}

	respondWithJSON(w, http.StatusOK, chirpRes)

}

func (cfg *apiConfig) getChirpHandler(w http.ResponseWriter, r *http.Request) {
	chirps, err := cfg.dbQueries.GetAllChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve chirps")
	}

	respondWithJSON(w, http.StatusOK, chirps)
}

func (cfg *apiConfig) getChirpByIDHandler(w http.ResponseWriter, r *http.Request) {
	chirpID := r.PathValue("chirpID")

	id, err := uuid.Parse(chirpID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID format")
		return
	}

	chirp, err := cfg.dbQueries.GetChirpByID(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Failed to retrieve chirp")
		return
	}

	respondWithJSON(w, http.StatusOK, chirp)
}

func (cfg *apiConfig) deleteChirpHandler(w http.ResponseWriter, r *http.Request) {

	tokenString, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	userID, err := auth.ValidateJWT(tokenString, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	chirpID := r.PathValue("chirpID")

	id, err := uuid.Parse(chirpID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID format")
		return
	}

	chirp, err := cfg.dbQueries.GetChirpByID(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Failed to retrieve chirp")
		return
	}

	if chirp.UserID != userID {
		respondWithError(w, http.StatusForbidden, "Unauthorized access")
		return
	}

	if err := cfg.dbQueries.DeleteChirp(r.Context(), chirp.ID); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
