package utils

// UnauthorizedError struct
type UnauthorizedError struct {
	message string
}

// NewUnauthorizedError func
func NewUnauthorizedError() *UnauthorizedError {
	return &UnauthorizedError{message: "401"}
}

func (e UnauthorizedError) Error() string {
	return e.message
}
