package error

import "errors"

var (
	InvalidPhoneNumber      = errors.New("Invalid phone number or password")
	TokenGenerateError      = errors.New("Failed to generate token")
	LoginInvalidJSONError   = errors.New("Invalid JSON input")
	ServiceError            = errors.New("Service not initialized")
	ErrInvalidPassword      = errors.New("Password must be at least 10 characters long and contain at least one uppercase letter, one lowercase letter, one number, and one special character")
	ErrUserNotFound         = errors.New("Password or Phonenumber is incorrect!s")
	ErrDbConnect            = errors.New("Can not get or tracking data")
	ErrLogin                = errors.New("Can not access login function! Please check it again")
	ErrCreatingToken        = errors.New("Something wrong in creating token for user")
	ErrInternalServer       = errors.New("Internal Server Error. Please check it again!")
	ErrBadRequest           = errors.New("Can not read the requests from client!")
	ErrInValidPhone         = errors.New("Invalid Phone Number! The phone number must start with 0, has at least 9 characters and maximum 11 characters")
	ErrEmptyLoginRequest    = errors.New("Phone or password are empty!")
	ErrEmptySignupRequest   = errors.New("Empty value")
	ErrDuplicatePhoneNumber = errors.New("Phone number is duplicated!")
	ErrInvalidEmail         = errors.New("Email is not valid")
	ErrDuplicateEmail       = errors.New("Email is duplicated")
	ErrInvalidToken         = errors.New("Invalid Token When Bearer Verify")
	ErrGetUserById          = errors.New("Failed in getting user by userId")
)
