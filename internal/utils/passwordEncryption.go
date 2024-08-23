package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a bcrypt hash for the given password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		fmt.Sprintf("Failed to hash the password : %w", err)
		return "", fmt.Errorf(" Error hashing the password ")
	}
	return string(bytes), err
}

// VerifyPassword verifies if the given password matches the stored hash.
func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
