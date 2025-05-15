package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (e *CustomError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}

func (e *CustomError) StatusCode() int {
	return e.Code
}

func (e *CustomError) Unwrap() error {
	return e.Err
}

func (e *CustomError) Respond(ctx *gin.Context) {
	ctx.JSON(e.Code, gin.H{
		"error":  e.Message,
		"status": e.Code,
	})
}

func NewBadRequestError(msg string) *CustomError {
	return &CustomError{Code: http.StatusBadRequest, Message: msg}
}

func NewNotFoundError(msg string) *CustomError {
	return &CustomError{Code: http.StatusNotFound, Message: msg}
}

func NewForbiddenError(msg string) *CustomError {
	return &CustomError{Code: http.StatusForbidden, Message: msg}
}

func NewInternalServerError(msg string, err error) *CustomError {
	return &CustomError{Code: http.StatusInternalServerError, Message: msg, Err: err}
}
