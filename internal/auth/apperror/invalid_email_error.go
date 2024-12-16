package apperror

import (
	"errors"

	"server/internal/auth/constant"
	"server/pkg/apperror"
)

func NewEmailNotExistsError() *apperror.AppError {
	msg := constant.EmailNotExistsErrorMessage

	err := errors.New(msg)

	return apperror.NewAppError(err, apperror.ForbiddenAccessErrorCode, msg)
}
