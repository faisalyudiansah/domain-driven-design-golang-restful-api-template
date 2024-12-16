package apperror

import (
	"errors"

	"server/internal/auth/constant"
	"server/pkg/apperror"
)

func NewTokenAlreadyExistsError() *apperror.AppError {
	msg := constant.TokenAlreadyExistsErrorMessage

	err := errors.New(msg)

	return apperror.NewAppError(err, apperror.TooManyRequestsErrorCode, msg)
}
