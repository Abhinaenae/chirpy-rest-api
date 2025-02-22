package main

import (
	"encoding/json"
	"net/http"
)

/*
type successRes struct {
	CleanedBody string `json:"cleaned_body"`
}*/

type errorRes struct {
	Error string `json:"error"`
}

func respondWithError(w http.ResponseWriter, status int, msg string) {
	respondWithJSON(w, status, errorRes{Error: msg})
}

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
