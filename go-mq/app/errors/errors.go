package errors

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Custom error types
type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

// Error constructors
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code:    fiber.StatusNotFound,
		Message: message,
		Err:     errors.New("not found"),
	}
}

func NewBadRequestError(message string) *AppError {
	return &AppError{
		Code:    fiber.StatusBadRequest,
		Message: message,
		Err:     errors.New("bad request"),
	}
}

func NewConflictError(message string) *AppError {
	return &AppError{
		Code:    fiber.StatusConflict,
		Message: message,
		Err:     errors.New("conflict"),
	}
}

func NewInternalError(err error) *AppError {
	return &AppError{
		Code:    fiber.StatusInternalServerError,
		Message: "Internal server error",
		Err:     errors.New("internal server error"),
	}
}

func NewUnauthorizedError(message string) *AppError {
	return &AppError{
		Code:    fiber.StatusNotFound,
		Message: message,
		Err:     errors.New("unauthorized"),
	}
}

func JWTError(message string, err error, code int) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
