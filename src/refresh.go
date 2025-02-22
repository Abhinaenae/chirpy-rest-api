package main

import (
	"chirpy/internal/auth"
	"database/sql"
	"errors"
	"net/http"
	"time"
)

type TokenResponse struct {
	Token string `json:"token"`
}

func (cfg *apiConfig) refreshHandler(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	userID, err := cfg.dbQueries.GetUserFromRefreshToken(r.Context(), refreshToken)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusUnauthorized, "invalid or expired refresh token")
			return
		}
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	newAccessToken, err := auth.MakeJWT(userID, cfg.jwtSecret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, TokenResponse{
		Token: newAccessToken,
	})
}
func (cfg *apiConfig) revokeHandler(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := cfg.dbQueries.RevokeRefreshToken(r.Context(), refreshToken); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
