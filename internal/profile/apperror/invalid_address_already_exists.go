package apperror

import (
	"errors"

	constantProfile "server/internal/profile/constant"
	"server/pkg/apperror"
)

func NewInvalidAddressAlreadyExistsError() *apperror.AppError {
	msg := constantProfile.InvalidAddressAlreadyExists
	err := errors.New(msg)
	return apperror.NewAppError(err, apperror.DefaultClientErrorCode, msg)
}
