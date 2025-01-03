package apperror

import (
	"errors"

	"server/internal/auth/constant"
	"server/pkg/apperror"
)

func NewInvalidEmailAlreadyExists(err error) *apperror.AppError {
	msg := constant.InvalidEmailAlreadyExists

	if err == nil {
		err = errors.New(msg)
	}

	return apperror.NewAppError(err, apperror.DefaultClientErrorCode, msg)
}

func NewInvalidSipaNumberAlreadyExists() *apperror.AppError {
	msg := constant.InvalidSipaNumberAlreadyExists
	err := errors.New(msg)
	return apperror.NewAppError(err, apperror.DefaultClientErrorCode, msg)
}

func NewInvalidWaNumberAlreadyExists() *apperror.AppError {
	msg := constant.InvalidWhatsappNumberAlreadyExists
	err := errors.New(msg)
	return apperror.NewAppError(err, apperror.DefaultClientErrorCode, msg)
}
