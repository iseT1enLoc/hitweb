package apperror

import (
	"errors"
	"fmt"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"statusCode" example:"400"`
	RootErr    error  `json:"-"`
	Message    string `json:"message" example:"error message"`
	Log        string `json:"log" example:"error log"`
	Key        string `json:"errorKey" example:"ErrKey"`
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewCustomError(root error, msg, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}

	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func (e *AppError) RootError() error {
	var err *AppError
	if errors.As(e.RootErr, &err) {
		return err.RootError()
	}
	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func ErrDB(err error) *AppError {
	return NewErrorResponse(
		err,
		"Something went wrong with database",
		err.Error(),
		"DB_ERROR",
	)
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(
		err,
		"Invalid request",
		err.Error(),
		"ERROR_INVALID_REQUEST",
	)
}

func ErrInternal(err error) *AppError {
	return NewErrorResponse(
		err,
		"Internal error",
		err.Error(),
		"ERROR_INVALID_REQUEST",
	)
}

func ErrNoPermission(err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf(err.Error()),
		fmt.Sprintf("ErrNoPermisstion"),
	)
}

func ErrRecordNotFound() *AppError {
	return NewCustomError(
		errRecordNotFound,
		fmt.Sprintf(errRecordNotFound.Error()),
		fmt.Sprintf("ErrRecordNotFound"),
	)
}

var (
	errRecordNotFound = errors.New("record not found")
)
