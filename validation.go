package utilities

import (
	"errors"
	"log"
	"regexp"

	"gopkg.in/validator.v2"
)

// ValidateRequiredFields validates if the required fields are present in the struct.
func ValidateRequiredFields(s interface{}) error {
	err := validator.Validate(s)
	if err != nil {
		log.Println(err.(validator.ErrorMap))
		return errors.New("missing fields")
	}
	return nil
}

// ValidateEmail validates if the email is valid.
func ValidateEmail(email string) error {
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !regex.MatchString(email) {
		return errors.New("invalid email")
	}
	return nil
}
