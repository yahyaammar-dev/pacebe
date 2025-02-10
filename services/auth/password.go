package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

// HashPassword hashes a password using SHA256.
func HashPassword(password string) (string, error) {
	if password == "" {
		return "", errors.New("password cannot be empty")
	}

	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:]), nil
}

// ComparePasswords compares a hashed password with a plain password.
func ComparePasswords(hashed string, plain []byte) bool {
	plainHash := sha256.Sum256(plain)
	return hashed == hex.EncodeToString(plainHash[:])
}
