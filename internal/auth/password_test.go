package auth

import (
	"testing"
)

// TestHashPassword ensures that hashing a password works and produces a valid hash.
func TestHashPassword(t *testing.T) {
	password := "securepassword"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	if len(hashedPassword) == 0 {
		t.Fatal("Hashed password is empty")
	}
}

// TestCheckPasswordHash verifies that a correct password matches the hash.
func TestCheckPasswordHash(t *testing.T) {
	password := "securepassword"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	// Test correct password
	if err := CheckPasswordHash(password, hashedPassword); err != nil {
		t.Errorf("Expected password to match, but got error: %v", err)
	}

	// Test incorrect password
	wrongPassword := "wrongpassword"
	if err := CheckPasswordHash(wrongPassword, hashedPassword); err == nil {
		t.Error("Expected error for incorrect password, but got nil")
	}
}
