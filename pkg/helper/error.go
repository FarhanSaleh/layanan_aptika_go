package helper

import (
	"fmt"

	"github.com/farhansaleh/layanan_aptika_be/internal/domain"
	"github.com/go-playground/validator/v10"
)
type(
	CustomValidationError struct {
		Errors []domain.ErrorsValidation
	}
	AuthError struct {
		Message string
	}
	BadRequestError struct {
		Message string
	}
)


func (e *CustomValidationError) Error() string {
	return fmt.Sprintf("validation failed with %d errors", len(e.Errors))
}

func MappingValidationError(err error) error {
	validatorError := err.(validator.ValidationErrors)
	errorResponse := []domain.ErrorsValidation{}
	for _, errorValidate := range validatorError {
		errorResponse = append(errorResponse, domain.ErrorsValidation{
			Field: errorValidate.Field(),
			Message: errorValidate.Error(),
		})
	}

	return &CustomValidationError{
		Errors: errorResponse,
	}
}

func IsValidationError(err error) (*CustomValidationError, bool) {
	validationErr, ok := err.(*CustomValidationError)
	return validationErr, ok
}

func NewAuthError(message string) *AuthError {
	return &AuthError{
		Message: message,
	}
}

func NewBadRequestError(message string) *BadRequestError {
	return &BadRequestError{
		Message: message,
	}
}

func (e *AuthError) Error() string {
	return e.Message
}

func (e *BadRequestError) Error() string {
	return e.Message
}
