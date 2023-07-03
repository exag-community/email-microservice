package utils

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"regexp"
)

// ValidateString validates a string
func ValidateString(value string, regex string) error {
	if !regexp.MustCompile(regex).MatchString(value) {
		return status.Errorf(codes.InvalidArgument, "The data provided is badly formatted")
	}

	return nil
}

func ValidateNonEmptyString(value string) error {
	if len(value) == 0 {
		log.Printf("value is empty: %v", value)
		return ErrorMessageFromStatusCode(&ErrorParams{
			Code:    codes.InvalidArgument,
			Message: "The ID provided is badly formatted",
		})
	}

	// check if string is not empty or blank
	if len(value) == 0 {
		log.Printf("value is empty: %v", value)
		return ErrorMessageFromStatusCode(&ErrorParams{
			Code:    codes.InvalidArgument,
			Message: "The provided value is empty or blank",
		})
	}

	return nil
}

// ValidateEmail validates an email
func ValidateEmail(email string) error {
	if err := ValidateString(email, `^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`); err != nil {
		log.Printf("email is invalid: %v", err)
		return ErrorMessageFromStatusCode(&ErrorParams{
			Code:    codes.InvalidArgument,
			Message: "Email is badly formatted",
		})
	}

	return nil
}
