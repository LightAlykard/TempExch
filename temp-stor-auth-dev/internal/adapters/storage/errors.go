package storage

import "errors"

var (
	NotFound   = errors.New("user with specified login was not found")
	UserExists = errors.New("given login is already in use")
	EmptyLogin = errors.New("login should not be empty")
	EmpthPass  = errors.New("pass should not be empty")
	NoSuchUser = errors.New("there is no user with such login")
)
