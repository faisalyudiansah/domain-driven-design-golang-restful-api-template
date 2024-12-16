package apperror

import (
	"errors"

	"server/internal/auth/constant"
	"server/pkg/apperror"
)

func NewUnverifiedError() *apperror.AppError {
	msg := constant.UnverifiedErrorMessage

	err := errors.New(msg)

	return apperror.NewAppError(err, apperror.ForbiddenAccessErrorCode, msg)
}
