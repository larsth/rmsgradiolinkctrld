package gps

import "errors"

var (
	ErrGpsdJsonDocumentHadNotBeenParsed = errors.New(
		"ERROR: The gpsd JSON document had not been parsed")
	//'bufferSize' is not a const, because its size is set to a much lower
	//value in unit tests
	ErrIoNilReader           = errors.New("Error: the io.Reader is nil")
	ErrBufferSizeLessThanOne = errors.New("The buffer size is less than 1")
	ErrNilBufioReader        = errors.New("Error: The buffered I/O Reader is nil")
)
