package v1

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckHashPassword(password, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hashPassword))
	return err == nil
}

func ValidatePassword(password string) bool {
	// Regular expression for validating password
	passwordRegex := regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*()-_=+]{8,}$`)

	return passwordRegex.MatchString(password)
}
