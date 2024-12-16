package apperror

import (
	"errors"

	"server/internal/profile/constant"
	"server/pkg/apperror"
)

func NewUserReachMaximumNumberOfAddress() *apperror.AppError {
	msg := constant.InvalidUserReachesMaximumNumberOfAddresses

	err := errors.New(msg)

	return apperror.NewAppError(err, apperror.DefaultClientErrorCode, msg)
}
