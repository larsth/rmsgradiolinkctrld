package gpspipe

import (
	"errors"
)

var (
	ErrEmptyDeviceName          = errors.New("Empty device name")
	ErrCmdHadBeenInitialized    = errors.New("The Cmd had been initialized")
	ErrCmdHadNotBeenInitialized = errors.New("The Cmd had NOT been initialized")
)
