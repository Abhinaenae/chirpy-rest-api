package main

import (
	"chirpy/internal/auth"
	"chirpy/internal/database"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

type newUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoggedInUser struct {
	ID           uuid.UUID `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Email        string    `json:"email"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
}

type userLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type userUpdateReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (cfg *apiConfig) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	decoder := json.NewDecoder(r.Body)
	updateUser := userUpdateReq{}
	if err := decoder.Decode(&updateUser); err != nil {
		log.Printf("Error decoding JSON: %s", err)
		respondWithError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	hashed_password, err := auth.HashPassword(updateUser.Password)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to parse password")
		return
	}

	user, err := cfg.dbQueries.UpdateUser(r.Context(), database.UpdateUserParams{
		Email:          updateUser.Email,
		HashedPassword: hashed_password,
		ID:             userID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Cannot store in Database")
		return
	}

	userRes := User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	}

	fmt.Printf("%v\n", userRes)

	respondWithJSON(w, http.StatusOK, userRes)
}

func (cfg *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	newUser := newUserRequest{}
	err := decoder.Decode(&newUser)
	if err != nil {
		// an error will be thrown if the JSON is invalid or has the wrong types
		// any missing fields will simply have their values in the struct set to their zero value
		log.Printf("Error decoding JSON: %s", err)
		respondWithError(w, 400, "Invalid JSON")
		return
	}

	hashed_password, err := auth.HashPassword(newUser.Password)
	if err != nil {
		respondWithError(w, 400, "Failed to parse password")
		return
	}

	user, err := cfg.dbQueries.CreateUser(r.Context(), database.CreateUserParams{
		Email:          newUser.Email,
		HashedPassword: hashed_password,
	})
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	userRes := User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	}

	fmt.Printf("%v\n", userRes)

	respondWithJSON(w, http.StatusCreated, userRes)
}

func (cfg *apiConfig) loginHandler(w http.ResponseWriter, r *http.Request) {
	//take the json of email and password
	decoder := json.NewDecoder(r.Body)
	loginReq := userLogin{}
	err := decoder.Decode(&loginReq)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	//look up the email in the db and see if the recieved password hash matches the db password hash
	foundUser, err := cfg.dbQueries.LookupUserbyEmail(r.Context(), loginReq.Email)
	if err != nil {
		//if either err, return 401 Unauthorized, error: incorrect email or password
		respondWithError(w, http.StatusUnauthorized, "incorrect email or password")
		return
	}

	//if either err, return 401 Unauthorized, error: incorrect email or password
	if err := auth.CheckPasswordHash(loginReq.Password, foundUser.HashedPassword); err != nil {
		respondWithError(w, http.StatusUnauthorized, "incorrect email or password")
		return
	}

	token, err := auth.MakeJWT(foundUser.ID, cfg.jwtSecret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error generating token")
		return
	}

	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error generating refresh token")
		return
	}

	refreshExpiration := time.Now().Add(60 * 24 * time.Hour)

	cfg.dbQueries.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{
		Token:     refreshToken,
		UserID:    foundUser.ID,
		ExpiresAt: refreshExpiration,
	})

	userRes := LoggedInUser{
		ID:           foundUser.ID,
		CreatedAt:    foundUser.CreatedAt,
		UpdatedAt:    foundUser.UpdatedAt,
		Email:        foundUser.Email,
		Token:        token,
		RefreshToken: refreshToken,
	}

	respondWithJSON(w, http.StatusOK, userRes)

}

func (cfg *apiConfig) resetHandler(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	err := cfg.dbQueries.DeleteAllUsers(r.Context())
	if err != nil {
		http.Error(w, "Failed to delete all users", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
