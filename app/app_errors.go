package app

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewErrorRespone(root error, message, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

func NewFullErrorRespone(statusCode int, root error, message, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorized(root error, message, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    message,
		Key:        key,
	}
}

func NewCustomError(root error, message, key string) *AppError {
	if root != nil {
		return NewErrorRespone(root, message, root.Error(), key)
	}
	return NewErrorRespone(errors.New(message), message, message, key)
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}
	return e.RootError()
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func ErrorDB(err error) *AppError {
	return NewErrorRespone(
		err,
		"Something went wrong with DB",
		err.Error(),
		"ErrorDB",
	)
}
func ErrorInvalidRequest(err error) *AppError {
	return NewErrorRespone(
		err,
		"Invalid request",
		err.Error(),
		"ErrorInvalidRequest",
	)
}

func ErrorInternal(err error) *AppError {
	return NewFullErrorRespone(
		http.StatusInternalServerError,
		err,
		"Something went wrong in the server",
		err.Error(),
		"ErrorInternal",
	)
}

func ErrorNoPermission(err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("You have no permission"),
		fmt.Sprintf("ErrorNoPermission"),
	)
}

func ErrorCanNotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Can not create %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrorCanNotCreate%s", entity),
	)
}

func ErrorCanNotFetchEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Can not fetch %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrorCanNotFetch%s", entity),
	)
}

func ErrorCanNotGetEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Can not get %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrorCanNotGet%s", entity),
	)
}

func ErrorCanNotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Can not update %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrorCanNotUpdate%s", entity),
	)
}

func ErrworCanNotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Can not delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrorCanNotDelete%s", entity),
	)
}

func ErrorEntityDeleted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s deleted", strings.ToLower(entity)),
		fmt.Sprintf("Error%sDeleted", entity),
	)
}

func ErrorEntityExisted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s already exists", strings.ToLower(entity)),
		fmt.Sprintf("Error%sAlreadyExists", entity),
	)
}

func ErrorEntityNotFound(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s not found", strings.ToLower(entity)),
		fmt.Sprintf("Error%sNotFound", entity),
	)
}
