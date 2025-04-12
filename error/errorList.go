package error

import "errors"

var (
	InvalidPhoneNumber    = errors.New("Invalid phone number or password")
	TokenGenerateError    = errors.New("Failed to generate token")
	LoginInvalidJSONError = errors.New("Invalid JSON input")
	ServiceError          = errors.New("Service not initialized")
)
