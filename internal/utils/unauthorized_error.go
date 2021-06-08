package utils

import (
	"net/http"
	"strconv"
)

// UnauthorizedError struct
type UnauthorizedError struct {
	message string
}

// NewUnauthorizedError func
func NewUnauthorizedError() *UnauthorizedError {
	return &UnauthorizedError{message: strconv.Itoa(http.StatusUnauthorized)}
}

func (e UnauthorizedError) Error() string {
	return e.message
}
