package auth_test

import (
	"chirpy/internal/auth"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestJWTCreationAndValidation(t *testing.T) {
	secret := "supersecretkey"
	userID := uuid.New()
	expiration := time.Minute * 5

	// Create JWT
	token, err := auth.MakeJWT(userID, secret, expiration)
	if err != nil {
		t.Fatalf("Failed to create JWT: %v", err)
	}

	// Validate JWT
	parsedUserID, err := auth.ValidateJWT(token, secret)
	if err != nil {
		t.Fatalf("Failed to validate JWT: %v", err)
	}

	// Check if user ID matches
	if parsedUserID != userID {
		t.Errorf("Expected user ID %v, got %v", userID, parsedUserID)
	}
}

func TestExpiredJWT(t *testing.T) {
	secret := "supersecretkey"
	userID := uuid.New()

	// Create an expired JWT
	token, err := auth.MakeJWT(userID, secret, -time.Minute)
	if err != nil {
		t.Fatalf("Failed to create JWT: %v", err)
	}

	// Validate JWT (should fail)
	_, err = auth.ValidateJWT(token, secret)
	if err == nil {
		t.Errorf("Expected error for expired JWT, but got none")
	}
}

func TestJWTWithWrongSecret(t *testing.T) {
	secret := "supersecretkey"
	wrongSecret := "wrongsecret"
	userID := uuid.New()
	expiration := time.Minute * 5

	// Create JWT
	token, err := auth.MakeJWT(userID, secret, expiration)
	if err != nil {
		t.Fatalf("Failed to create JWT: %v", err)
	}

	// Validate JWT with wrong secret (should fail)
	_, err = auth.ValidateJWT(token, wrongSecret)
	if err == nil {
		t.Errorf("Expected error for JWT with wrong secret, but got none")
	}
}
func TestGetBearerToken(t *testing.T) {
	tests := []struct {
		name      string
		headers   http.Header
		wantToken string
		wantErr   bool
	}{
		{
			name: "Valid Bearer Token",
			headers: http.Header{
				"Authorization": []string{"Bearer valid-token"},
			},
			wantToken: "valid-token",
			wantErr:   false,
		},
		{
			name:    "Missing Authorization Header",
			headers: http.Header{},
			wantErr: true,
		},
		{
			name: "Invalid Format (No Bearer Prefix)",
			headers: http.Header{
				"Authorization": []string{"invalid-token"},
			},
			wantErr: true,
		},
		{
			name: "Empty Token After Bearer Prefix",
			headers: http.Header{
				"Authorization": []string{"Bearer "},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := auth.GetBearerToken(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("expected error: %v, got: %v", tt.wantErr, err)
			}
			if token != tt.wantToken {
				t.Errorf("expected token: %s, got: %s", tt.wantToken, token)
			}
		})
	}
}
