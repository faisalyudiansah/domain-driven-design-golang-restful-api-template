package apperror

import (
	"errors"
	"fmt"

	"server/internal/profile/constant"
	"server/pkg/apperror"
)

func NewInvalidImageFormatErrorMessageError() *apperror.AppError {
	msg := fmt.Sprintf(constant.InvalidImageFormatErrorMessage)
	err := errors.New(msg)
	return apperror.NewAppError(err, apperror.DefaultClientErrorCode, msg)
}

func NewInvalidPhotoMaxSizeError() *apperror.AppError {
	msg := fmt.Sprintf(constant.InvalidPhotoMaxSize, fmt.Sprintf("%vkb", constant.MAX_IMAGE_SIZE/1024))
	err := errors.New(msg)
	return apperror.NewAppError(err, apperror.DefaultClientErrorCode, msg)
}
