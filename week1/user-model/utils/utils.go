package utils

import (
	"net/mail"

	"github.com/google/uuid"
)

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)

	if err != nil {
		return false
	} else {
		return true
	}
}

func GenerateUUID() string {
	return uuid.NewString()
}
