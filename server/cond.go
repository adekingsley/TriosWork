package server

import (
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

// / containsAlphabetAndNumber checks if a string contains both alphabet characters and numbers
func containsAlphabetAndNumber(s string) bool {
	hasAlphabet := false
	hasNumber := false

	for _, char := range s {
		if unicode.IsLetter(char) {
			hasAlphabet = true
		} else if unicode.IsNumber(char) {
			hasNumber = true
		}
	}

	// If both alphabet characters and numbers are found, return true
	return hasAlphabet && hasNumber
}

// hashPassword hashes the given password using bcrypt
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
