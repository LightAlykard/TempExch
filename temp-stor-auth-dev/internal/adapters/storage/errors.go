package storage

import "errors"

var (
	NOT_FOUND = errors.New("user with specified login was not found")
)
