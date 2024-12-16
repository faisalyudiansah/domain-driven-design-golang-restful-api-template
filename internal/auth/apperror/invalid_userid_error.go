package apperror

import (
	"errors"

	"server/internal/auth/constant"
	"server/pkg/apperror"
)

func NewInvalidUserIdError() *apperror.AppError {
	msg := constant.InvalidUserId

	err := errors.New(msg)

	return apperror.NewAppError(err, apperror.DefaultClientErrorCode, msg)
}
