package errs

import "errors"

var (
	ErrKeyNotFound = errors.New("environment variable not found")
	ErrMinValue    = errors.New("value is less than minimum")
	ErrMaxValue    = errors.New("value is greater than maximum")
	ErrNotInteger  = errors.New("value is not an integer")
)
