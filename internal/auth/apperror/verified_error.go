package apperror

import (
	"errors"

	"server/internal/auth/constant"
	"server/pkg/apperror"
)

func NewVerifiedError() *apperror.AppError {
	msg := constant.AlreadyVerifiedErrorMessage

	err := errors.New(msg)

	return apperror.NewAppError(err, apperror.DefaultClientErrorCode, msg)
}
