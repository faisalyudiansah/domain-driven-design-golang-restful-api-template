package apperror

import (
	"errors"

	"server/internal/auth/constant"
	"server/pkg/apperror"
)

func NewInvalidLoginCredentials(err error) *apperror.AppError {
	msg := constant.InvalidLoginCredentialsErrorMessage

	if err == nil {
		err = errors.New(msg)
	}

	return apperror.NewAppError(err, apperror.DefaultClientErrorCode, msg)
}
