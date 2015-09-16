package gps

import (
	"errors"
	"io"
)

const DefaultBufferSize = 8192

var (
	ErrGpsdJsonDocumentHadNotBeenParsed = errors.New(
		"ERROR: The gpsd JSON document had not been parsed")
	//'bufferSize' is not a const, because its size is set to a much lower
	//value in unit tests
	defaultBufferSize        = 4096
	ErrIoNilReader           = errors.New("Error: the io.Reader is nil")
	ErrBufferSizeLessThanOne = errors.New("The buffer size is less than 1")
)

func (coord *GPSCoord) ReadFromIOReader(r io.Reader) error {
	coord.mutex.Lock()
	defer coord.mutex.Unlock()

	if r == nil {
		return ErrIoNilReader
	}
	if r != coord.ioReader {
		coord.ioReader = r
		coord.scanBuf = make([]byte, 0, DefaultBufferSize)
	}

	return nil
}
