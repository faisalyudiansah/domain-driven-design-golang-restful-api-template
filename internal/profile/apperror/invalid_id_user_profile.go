package apperror

import (
	"errors"

	constantProfile "server/internal/profile/constant"
	"server/pkg/apperror"
)

func NewInvalidIdUserProfileError() *apperror.AppError {
	msg := constantProfile.InvalidIdUserProfile
	err := errors.New(msg)
	return apperror.NewAppError(err, apperror.DefaultClientErrorCode, msg)
}

func NewInvalidIdUserProfileNotExistsError() *apperror.AppError {
	msg := constantProfile.InvalidIdUserProfileDoesNotExists
	err := errors.New(msg)
	return apperror.NewAppError(err, apperror.DefaultClientErrorCode, msg)
}
