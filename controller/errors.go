package controller

import "errors"

var (
	ErrMissingField       = errors.New("missing field")
	ErrInvalidRequestBody = errors.New("invalid request body")
)
