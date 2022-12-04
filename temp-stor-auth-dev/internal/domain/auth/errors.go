package auth

import "errors"

var (
	ErrPassIncorrect = errors.New("password mismatch")
	ErrEmptySecret   = errors.New("empty secret key for token creation")
	ErrZeroDuration  = errors.New("zero duration for token creation")
	ErrNotFound      = errors.New("user with specified login was not found")
	ErrUserExists    = errors.New("given login is already in use")
	ErrEmptyLogin    = errors.New("login should not be empty")
	ErrEmpthPass     = errors.New("pass should not be empty")
	ErrNoSuchUser    = errors.New("there is no user with such login")
)
