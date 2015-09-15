package errs

import (
	"errors"
)

var (
	ErrEmptyString       = errors.New("ERROR: The empty string")
	ErrZeroLengthPayload = errors.New("ERROR: `payload` has no octets (length is zero)")
)
