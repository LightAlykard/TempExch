package storage

import "errors"

var (
	NotFound = errors.New("user with specified login was not found")
)
