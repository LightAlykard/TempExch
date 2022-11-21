package auth

import "errors"

var (
	PASS_INCORRECT = errors.New("password mismatch")
	EMPTY_SECRET   = errors.New("empty secret key for token creation")
	ZERO_DURATION  = errors.New("zero duration for token creation")
)
