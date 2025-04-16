package domain

import "errors"

var (
	// ErrAccountNotFound is returned when an account is not found
	ErrAccountNotFound = errors.New("account not found")
	// ErrDuplicatedAPIKey is returned when a duplicated API key is used
	ErrDuplicatedAPIKey = errors.New("api key already exists")
	// ErrInvoiceNotFound is returned when an invoice is not found
	ErrInvoiceNotFound = errors.New("invoice not found")
	// ErrUnauthorizedAccess is returned when an unauthorized access is attempted
	ErrUnauthorizedAccess = errors.New("unauthorized access")
	// ErrInvalidAmount is returned when an invalid amount is used
	ErrInvalidAmount = errors.New("invalid amount")
	// ErrInvalidStatus is returned when an invalid status is used
	ErrInvalidStatus = errors.New("invalid status")
)
