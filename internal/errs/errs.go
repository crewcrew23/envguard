package errs

import (
	"errors"
)

var (
	ErrKeyNotFound = errors.New("environment variable not found")
	ErrMinValue    = errors.New("value is less than minimum")
	ErrMaxValue    = errors.New("value is greater than maximum")
	ErrNotInteger  = errors.New("value is not an integer")
	ErrNotEven     = errors.New("value is not a even")
	ErrNotOdd      = errors.New("value is not a odd")
	ErrNotPositive = errors.New("value is not positive")
	ErrNotNegative = errors.New("value is not negative")
	ErrZero        = errors.New("value is zero")

	ErrStrMin      = errors.New("length is less than minimum")
	ErrStrMax      = errors.New("length is greater than minimum")
	ErrStrEmpty    = errors.New("string is empty")
	ErrStrLength   = errors.New("string length is incorrect")
	ErrStrBlank    = errors.New("string is blank")
	ErrNotEmail    = errors.New("is not email")
	ErrRegexp      = errors.New("custom regexp is failed")
	ErrNotUrl      = errors.New("is not URL")
	ErrNotContains = errors.New("value not contains in string")
	ErrContains    = errors.New("value contains in string")
)
