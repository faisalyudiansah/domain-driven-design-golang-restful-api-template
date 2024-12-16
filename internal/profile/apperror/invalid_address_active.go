package apperror

import (
	"errors"

	constantProfile "server/internal/profile/constant"
	"server/pkg/apperror"
)

func NewInvalidAddressActiveNotExistsError() *apperror.AppError {
	msg := constantProfile.InvalidAddressActiveNotExists
	err := errors.New(msg)
	return apperror.NewAppError(err, apperror.DefaultClientErrorCode, msg)
}
