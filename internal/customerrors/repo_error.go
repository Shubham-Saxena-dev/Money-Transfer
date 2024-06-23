package customerrors

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidAcc          = errors.New("account does not exit")
	ErrSave                = errors.New("failed to save in repo")
	ErrUpdate              = errors.New("failed to update from repo")
	ErrInsufficientBalance = errors.New("insufficient balance")
)

type RepositoryErrors struct {
	errType error
	cause   error
}

func newRepositoryErrors(cause, errType error) *RepositoryErrors {
	return &RepositoryErrors{errType: errType, cause: cause}
}

func (ve *RepositoryErrors) Error() string {
	return fmt.Sprintf("%v: %v", ve.cause, ve.errType)
}

func (ve *RepositoryErrors) Unwrap() error {
	return ve.errType
}

func ErrAcc(cause error) *RepositoryErrors {
	return newRepositoryErrors(cause, ErrInvalidAcc)
}

func ErrSaveRepo(cause error) *RepositoryErrors {
	return newRepositoryErrors(cause, ErrSave)
}

func ErrUpdateRepo(cause error) *RepositoryErrors {
	return newRepositoryErrors(cause, ErrUpdate)
}

func ErrBalance() *RepositoryErrors {
	return newRepositoryErrors(nil, ErrInsufficientBalance)
}
