package auth

import "errors"

var (
	PassIncorrect = errors.New("password mismatch")
	EmptySecret   = errors.New("empty secret key for token creation")
	ZeroDuration  = errors.New("zero duration for token creation")
)
