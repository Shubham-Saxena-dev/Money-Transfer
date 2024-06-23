package customerrors

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidOrgBIC   = errors.New("invalid org BIC")
	ErrInvalidOrgIBAN  = errors.New("invalid org IBAN")
	ErrInvalidTransfer = errors.New("invalid transfer")
	ErrInvalidValue    = errors.New("invalid value")
)

type ValidationErrors struct {
	key     string
	errType error
}

func newValidationErrors(key string, errType error) *ValidationErrors {
	return &ValidationErrors{errType: errType, key: key}
}

func (ve *ValidationErrors) Error() string {
	return fmt.Sprintf("%s: %v", ve.key, ve.errType)
}

func (ve *ValidationErrors) Unwrap() error {
	return ve.errType
}

func ErrOrgBIC(key string) *ValidationErrors {
	return newValidationErrors(key, ErrInvalidOrgBIC)
}

func ErrOrgIBAN(key string) *ValidationErrors {
	return newValidationErrors(key, ErrInvalidOrgIBAN)
}

func ErrTransfer(key string) *ValidationErrors {
	return newValidationErrors(key, ErrInvalidTransfer)
}

func ErrValue(key string) *ValidationErrors {
	return newValidationErrors(key, ErrInvalidValue)
}
