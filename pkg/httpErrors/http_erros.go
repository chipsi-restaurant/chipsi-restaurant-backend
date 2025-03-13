package httpErrors

import (
	"errors"
	"fmt"
)

var (
	BadRequest          = errors.New("Bad Request")
	Unauthorized        = errors.New("Unauthorized")
	InvalidJWTToken     = errors.New("Invalid JWT token")
	ErrNoCookie         = errors.New("No cookie")
	InvalidJWTClaims    = errors.New("Invalid JWT claims")
	NotFound            = errors.New("Not Found")
	InternalServerError = errors.New("Internal Server Error")
	ExistsEmailError    = errors.New("User with given email already exists")
)

type RestErr interface {
	Status() int
	Error() string
	Causes() interface{}
}

type RestError struct {
	ErrStatus int         `json:"status,omitempty"`
	ErrError  string      `json:"error,omitempty"`
	ErrCauses interface{} `json:"-"`
}

func (e RestError) Error() string {
	return fmt.Sprintf("status: %d - errors: %s - causes: %v", e.ErrStatus, e.ErrError, e.ErrCauses)
}

func (e RestError) Status() int {
	return e.ErrStatus
}

func (e RestError) Causes() interface{} {
	return e.ErrCauses
}

func NewRestError(status int, err string, causes interface{}) RestErr {
	return RestError{
		ErrStatus: status,
		ErrError:  err,
		ErrCauses: causes,
	}
}
