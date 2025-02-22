package main

import (
	"chirpy/internal/database"
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	dbQueries      *database.Queries
	platform       string
	jwtSecret      string
}

func main() {
	godotenv.Load()
	jwtSecret := os.Getenv("JWT_SECRET")
	platform := os.Getenv("PLATFORM")
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("DB Connection failed: %v", err) // Terminates the program

	}

	dbQueries := database.New(db)

	apiCfg := apiConfig{
		dbQueries:      dbQueries,
		fileserverHits: atomic.Int32{},
		platform:       platform,
		jwtSecret:      jwtSecret,
	}
	fs := http.FileServer(http.Dir("."))

	mux := http.NewServeMux()

	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", fs)))

	mux.HandleFunc("GET /api/healthz", readinessHandler)

	mux.HandleFunc("POST /admin/reset", apiCfg.resetHandler)
	mux.HandleFunc("GET /admin/metrics", apiCfg.metricsHandler)

	mux.HandleFunc("POST /api/users", apiCfg.createUserHandler)
	mux.HandleFunc("POST /api/login", apiCfg.loginHandler)
	mux.HandleFunc("PUT /api/users", apiCfg.updateUserHandler)

	mux.HandleFunc("POST /api/chirps", apiCfg.postChirpHandler)
	mux.HandleFunc("GET /api/chirps", apiCfg.getChirpHandler)
	mux.HandleFunc("GET /api/chirps/{chirpID}", apiCfg.getChirpByIDHandler)
	mux.HandleFunc("DELETE /api/chirps/{chirpID}", apiCfg.deleteChirpHandler)

	mux.HandleFunc("POST /api/refresh", apiCfg.refreshHandler)
	mux.HandleFunc("POST /api/revoke", apiCfg.revokeHandler)

	server := &http.Server{
		Addr:    ":8082",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())

}
